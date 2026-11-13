package dockerfile

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content/code"
)

// scan returns the spans of the document in the order they are emitted, each
// as its line, its text and its token, checking they are within the lines.
func scan(t *testing.T, src string) []string {
	t.Helper()
	lines := strings.Split(src, "\n")
	var got []string
	Scan(lines, func(s Span) {
		if s.Line < 0 || s.Line >= len(lines) || s.Start < 0 || s.End > len(lines[s.Line]) || s.Start >= s.End {
			t.Fatalf("span %+v is out of range", s)
		}
		got = append(got, fmt.Sprintf("%d %q %s", s.Line, lines[s.Line][s.Start:s.End], s.Token))
	})
	return got
}

func TestScan(t *testing.T) {
	for _, tc := range []struct {
		name, src string
		want      []string
	}{
		{
			// The instructions, in any case, and the image and the name of a stage. ONBUILD takes another instruction, but not FROM, and MAINTAINER expands no variables.
			name: "instructions",
			src:  "FROM --platform=$BUILDPLATFORM golang:${GO_VERSION:-1.23} AS build\nfrom alpine as Final\nWORKDIR /src\nUSER app:app\nonbuild run make\nONBUILD FROM x\nNOPE x\nMAINTAINER Me <me@$HOST>\nSTOPSIGNAL SIGTERM\nVOLUME [\"/data\", \"$D\"]",
			want: []string{
				"0 \"FROM\" Keyword",
				"0 \"$\" Escape",
				"0 \"BUILDPLATFORM\" Field",
				"0 \"--platform\" Constant",
				"0 \"golang:${GO_VERSION:-1.23}\" ImportRef",
				"0 \"${\" Escape",
				"0 \"GO_VERSION\" Field",
				"0 \"}\" Escape",
				"0 \"AS\" Keyword",
				"0 \"build\" FuncDeclaration",
				"1 \"from\" Keyword",
				"1 \"alpine\" ImportRef",
				"1 \"as\" Keyword",
				"1 \"Final\" FuncDeclaration",
				"2 \"WORKDIR\" Keyword",
				"3 \"USER\" Keyword",
				"4 \"onbuild\" Keyword",
				"4 \"run\" Keyword",
				"5 \"ONBUILD\" Keyword",
				"7 \"MAINTAINER\" Keyword",
				"8 \"STOPSIGNAL\" Keyword",
				"9 \"VOLUME\" Keyword",
				"9 \"\\\"/data\\\"\" StringLiteral",
				"9 \"\\\"$D\\\"\" StringLiteral",
				"9 \"$\" Escape",
				"9 \"D\" Field",
			},
		},
		{
			// The parser directives at the top, in any case and with spaces around =, until the first line that is not a known directive.
			name: "directives",
			src:  "# syntax=docker/dockerfile:1\n#ESCAPE = `\n  # check=skip=JSONArgsRecommended;error=true\n# unknown=x\n# syntax=again\nFROM x",
			want: []string{
				"0 \"# syntax=docker/dockerfile:1\" Directive",
				"1 \"#ESCAPE = `\" Directive",
				"2 \"# check=skip=JSONArgsRecommended;error=true\" Directive",
				"3 \"# unknown=x\" Comment",
				"4 \"# syntax=again\" Comment",
				"5 \"FROM\" Keyword",
				"5 \"x\" ImportRef",
			},
		},
		{
			// An empty line ends the directives: the escape one after it is a comment, so \ still continues the lines.
			name: "directive ends",
			src:  "# syntax=docker/dockerfile:1\n\n# escape=`\nRUN a \\\n b",
			want: []string{
				"0 \"# syntax=docker/dockerfile:1\" Directive",
				"2 \"# escape=`\" Comment",
				"3 \"RUN\" Keyword",
			},
		},
		{
			// A repeated directive fails the build, and is not taken.
			name: "repeated directive",
			src:  "# syntax=a\n# syntax=b\nFROM x",
			want: []string{
				"0 \"# syntax=a\" Directive",
				"1 \"# syntax=b\" Comment",
				"2 \"FROM\" Keyword",
				"2 \"x\" ImportRef",
			},
		},
		{
			// An escape character other than ` and \ fails the build, and is not taken.
			name: "invalid escape",
			src:  "# escape=x\nRUN a \\\n b",
			want: []string{
				"0 \"# escape=x\" Comment",
				"1 \"RUN\" Keyword",
			},
		},
		{
			// After a shebang, the syntax directive is still taken by the detection of the frontend, with no space before #, while the others are comments.
			name: "shebang",
			src:  "#!/usr/bin/env buildctl\n# escape=`\n# syntax=docker/dockerfile:1\n  # syntax=ignored\nFROM x",
			want: []string{
				"0 \"#!/usr/bin/env buildctl\" Directive",
				"2 \"# syntax=docker/dockerfile:1\" Directive",
				"1 \"# escape=`\" Comment",
				"3 \"# syntax=ignored\" Comment",
				"4 \"FROM\" Keyword",
				"4 \"x\" ImportRef",
			},
		},
		{
			// The byte order mark is not a part of the first line.
			name: "bom",
			src:  "\uFEFF# syntax=x\nFROM y",
			want: []string{
				"0 \"# syntax=x\" Directive",
				"1 \"FROM\" Keyword",
				"1 \"y\" ImportRef",
			},
		},
		{
			// The escape directive changes the character continuing the lines and escaping the others, like $.
			name: "backtick escape",
			src:  "# escape=`\nRUN dir C:\\ `\n  && echo `$X $Y\nCOPY a\\ b",
			want: []string{
				"0 \"# escape=`\" Directive",
				"1 \"RUN\" Keyword",
				"2 \"$\" Escape",
				"2 \"Y\" Field",
				"3 \"COPY\" Keyword",
			},
		},
		{
			// The comments and the empty lines between the continued lines are skipped, an escaped escape character continues nothing, and a string may continue on the next line.
			name: "continuations",
			src:  "RUN apt-get update \\\n  # a comment\n\n  && apt-get install -y \\\n     curl\nRUN echo \\\\\nRUN \"a\\\n b\" \\\r\n  'c'\r\nCMD x",
			want: []string{
				"1 \"# a comment\" Comment",
				"0 \"RUN\" Keyword",
				"5 \"RUN\" Keyword",
				"6 \"RUN\" Keyword",
				"6 \"\\\"a\" StringLiteral",
				"7 \" b\\\"\" StringLiteral",
				"8 \"'c'\" StringLiteral",
				"9 \"CMD\" Keyword",
			},
		},
		{
			// The variables are expanded in double quotes, but not in single ones or escaped. The values of ARG, ENV and LABEL follow their names, in the legacy form of ENV too.
			name: "strings and vars",
			src:  "RUN echo \"$HOME ${A:-${B:-$C}} \\$X\" '$Y' $ ${unclosed $Z\nENV PATH=\"/go/bin:$PATH\" EMPTY= Q='$q'\nENV LEGACY value with $V\nLABEL \"com.example.vendor\"=\"ACME\" version=1.0\nARG A B=${A}\nEXPOSE 80/tcp 8000-8010 $PORT",
			want: []string{
				"0 \"RUN\" Keyword",
				"0 \"\\\"$HOME ${A:-${B:-$C}} \\\\$X\\\"\" StringLiteral",
				"0 \"$\" Escape",
				"0 \"HOME\" Field",
				"0 \"${\" Escape",
				"0 \"A\" Field",
				"0 \"${\" Escape",
				"0 \"B\" Field",
				"0 \"$\" Escape",
				"0 \"C\" Field",
				"0 \"}\" Escape",
				"0 \"}\" Escape",
				"0 \"'$Y'\" StringLiteral",
				"0 \"${\" Escape",
				"0 \"unclosed\" Field",
				"0 \"$\" Escape",
				"0 \"Z\" Field",
				"1 \"ENV\" Keyword",
				"1 \"PATH\" Field",
				"1 \"\\\"/go/bin:$PATH\\\"\" StringLiteral",
				"1 \"$\" Escape",
				"1 \"PATH\" Field",
				"1 \"EMPTY\" Field",
				"1 \"Q\" Field",
				"1 \"'$q'\" StringLiteral",
				"2 \"ENV\" Keyword",
				"2 \"LEGACY\" Field",
				"2 \"$\" Escape",
				"2 \"V\" Field",
				"3 \"LABEL\" Keyword",
				"3 \"\\\"com.example.vendor\\\"\" Field",
				"3 \"\\\"ACME\\\"\" StringLiteral",
				"3 \"version\" Field",
				"4 \"ARG\" Keyword",
				"4 \"A\" Field",
				"4 \"B\" Field",
				"4 \"${\" Escape",
				"4 \"A\" Field",
				"4 \"}\" Escape",
				"5 \"EXPOSE\" Keyword",
				"5 \"80\" NumberLiteral",
				"5 \"8000-8010\" NumberLiteral",
				"5 \"$\" Escape",
				"5 \"PORT\" Field",
			},
		},
		{
			// The exec form, a JSON array, expands no variables of a command, while COPY expands them in both forms. An invalid array is taken in the shell form.
			name: "json",
			src:  "CMD [\"echo\", \"$HOME\", \"a\\\"b\\u00e9\"]\nENTRYPOINT [\"bad\" \"json\"]\nCOPY [\"$SRC\", \"/dst\"]\nSHELL [\"pwsh\", \"-c\"]",
			want: []string{
				"0 \"CMD\" Keyword",
				"0 \"\\\"echo\\\"\" StringLiteral",
				"0 \"\\\"$HOME\\\"\" StringLiteral",
				"0 \"\\\"a\\\\\\\"b\\\\u00e9\\\"\" StringLiteral",
				"0 \"\\\\\\\"\" Escape",
				"0 \"\\\\u00e9\" Escape",
				"1 \"ENTRYPOINT\" Keyword",
				"1 \"\\\"bad\\\"\" StringLiteral",
				"1 \"\\\"json\\\"\" StringLiteral",
				"2 \"COPY\" Keyword",
				"2 \"\\\"$SRC\\\"\" StringLiteral",
				"2 \"$\" Escape",
				"2 \"SRC\" Field",
				"2 \"\\\"/dst\\\"\" StringLiteral",
				"3 \"SHELL\" Keyword",
				"3 \"\\\"pwsh\\\"\" StringLiteral",
				"3 \"\\\"-c\\\"\" StringLiteral",
			},
		},
		{
			// The heredocs of RUN, COPY and ADD, with ONBUILD too, follow the instruction. A quoted word expands nothing in the body, and <<- ignores the tabs before the terminator.
			name: "heredocs",
			src:  "RUN <<EOF bash\n  echo $HOME\nEOF\nCOPY <<-\"END\" 2<<'X' /out\n\techo $NOT\n\tEND\nsecond $NOT\nX\nonbuild run cat <<A\n${V}\nA\nCMD <<EOF\nFROM x",
			want: []string{
				"0 \"RUN\" Keyword",
				"0 \"<<EOF\" StringLiteral",
				"1 \"  echo $HOME\" StringLiteral",
				"1 \"$\" Escape",
				"1 \"HOME\" Field",
				"2 \"EOF\" StringLiteral",
				"3 \"COPY\" Keyword",
				"3 \"\\\"END\\\"\" StringLiteral",
				"3 \"'X'\" StringLiteral",
				"3 \"<<-\\\"END\\\"\" StringLiteral",
				"3 \"2<<'X'\" StringLiteral",
				"4 \"\\techo $NOT\" StringLiteral",
				"5 \"\\tEND\" StringLiteral",
				"6 \"second $NOT\" StringLiteral",
				"7 \"X\" StringLiteral",
				"8 \"onbuild\" Keyword",
				"8 \"run\" Keyword",
				"8 \"<<A\" StringLiteral",
				"9 \"${V}\" StringLiteral",
				"9 \"${\" Escape",
				"9 \"V\" Field",
				"9 \"}\" Escape",
				"10 \"A\" StringLiteral",
				"11 \"CMD\" Keyword",
				"12 \"FROM\" Keyword",
				"12 \"x\" ImportRef",
			},
		},
		{
			// An unterminated heredoc runs to the end of the document.
			name: "heredoc unterminated",
			src:  "RUN <<EOF\nline\nFROM x",
			want: []string{
				"0 \"RUN\" Keyword",
				"0 \"<<EOF\" StringLiteral",
				"1 \"line\" StringLiteral",
				"2 \"FROM x\" StringLiteral",
			},
		},
		{
			// HEALTHCHECK runs a command after its flags, or takes NONE.
			name: "healthcheck",
			src:  "HEALTHCHECK --interval=30s CMD [\"curl\", \"-f\"]\nHEALTHCHECK NONE\nHEALTHCHECK CMD curl $URL",
			want: []string{
				"0 \"HEALTHCHECK\" Keyword",
				"0 \"--interval\" Constant",
				"0 \"CMD\" Keyword",
				"0 \"\\\"curl\\\"\" StringLiteral",
				"0 \"\\\"-f\\\"\" StringLiteral",
				"1 \"HEALTHCHECK\" Keyword",
				"1 \"NONE\" Keyword",
				"2 \"HEALTHCHECK\" Keyword",
				"2 \"CMD\" Keyword",
				"2 \"$\" Escape",
				"2 \"URL\" Field",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := scan(t, tc.src)
			if strings.Join(got, "\n") != strings.Join(tc.want, "\n") {
				t.Errorf("got\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(tc.want, "\n"))
			}
		})
	}
}

func TestScan_Linear(t *testing.T) {
	// The constructs with no end on a long line are scanned in linear time, and so
	// are the nested ones. Each of these takes seconds when the rest of the line
	// is scanned again from each opener, or for each nesting level.
	const n = 100_000
	r := strings.Repeat
	for name, src := range map[string]string{
		"unclosed variables": "RUN " + r("${a ", n),
		"nested variables":   "RUN " + r("${a:-", n) + r("}", n),
		"unclosed strings":   "RUN " + r("\"'", n),
		"continued lines":    r("RUN a \\\n", n) + "b",
		"heredocs":           "RUN " + r("<<A ", n) + "\nA",
		"many words":         "ENV " + r("a=$b ", n),
		"json":               "CMD [" + r(`"$a\n", `, n) + `"b"]`,
	} {
		painted := 0
		start := time.Now()
		Scan(strings.Split(src, "\n"), func(s Span) { painted += s.End - s.Start })
		if d := time.Since(start); d > time.Second {
			t.Errorf("%s: scanning %d bytes takes %v", name, len(src), d)
		}
		// The spans are painted for the editor, a nested span over its parent.
		if painted > 8*len(src) {
			t.Errorf("%s: the spans of %d bytes cover %d", name, len(src), painted)
		}
	}
}

func TestScan_Incomplete(t *testing.T) {
	// The instructions cut short, as they are while typed, end on their line.
	for _, src := range []string{"#", "FROM", "RUN <<", `RUN <<""`, "ENV", "LABEL =", "EXPOSE -", "CMD [", "RUN ${", `RUN "$`, "HEALTHCHECK", "ONBUILD"} {
		var tokens []code.TokenType
		Scan([]string{src, "# done"}, func(s Span) {
			if s.Line == 1 {
				tokens = append(tokens, s.Token)
			}
		})
		if len(tokens) != 1 || tokens[0] != code.TtComment {
			t.Errorf("%q: the next line is %v, want a comment", src, tokens)
		}
	}
}
