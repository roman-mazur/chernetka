package extsyntaxhl

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/styles"
	"rmazur.io/chernetka/internal/vt/escape"
)

// TestVisual renders highlighted buffers into the test log so the colors can be
// inspected by eye with `go test -v ./internal/editor/extsyntaxhl/`. It asserts
// nothing beyond the buffer being highlighted at all: it is here to be looked
// at, not to catch regressions. The exact offsets are covered by the tests in
// markdown_test.go and treesitter_test.go.
func TestVisual(t *testing.T) {
	t.Log("token colors:\n" + tokenLegend())

	for _, tc := range []struct{ path, text string }{
		{"visual.md", visualMarkdown},
		{"visual.go", visualGo},
		{"visual.nix", visualNix},
		{"visual.cue", visualCUE},
		{"visual.json", visualJSON},
		{"visual.yaml", visualYAML},
		{"visual.sh", visualShell},
		{"visual.c", visualC},
		{"visual.cpp", visualCPP},
		{"visual.cs", visualCSharp},
		{".git/COMMIT_EDITMSG", visualGitMsg},
		{".git/rebase-merge/git-rebase-todo", visualGitRebase},
	} {
		t.Run(tc.path, func(t *testing.T) {
			h := editor.NewTestHarness(t)
			h.Extend(new(Integration))
			if err := h.OpenReader(tc.path, strings.NewReader(tc.text)); err != nil {
				t.Fatal(err)
			}
			buf := h.Top()

			hl, ok := buf.ExtensionData("syntaxhl").(editor.SyntaxHighlighter)
			if !ok {
				t.Fatalf("no syntax highlighter for %q", tc.path)
			}

			t.Log("rendered:\n" + h.RenderBuffer(100))

			// The same content again, annotated with the token type behind every
			// colorized run, for when the colors alone are ambiguous.
			var spans strings.Builder
			for i, line := range buf.Content.Lines() {
				text := line.String()
				names := make([]string, 0, 8)
				for _, s := range hl.SyntaxSpans(i, text) {
					names = append(names, fmt.Sprintf("%s=%q", s.TokenType, text[s.Start:s.End]))
				}
				fmt.Fprintf(&spans, "%2d | %-52s | %s\n", i+1, text, strings.Join(names, " "))
			}
			t.Log("spans:\n" + spans.String())
		})
	}
}

func tokenLegend() string {
	var out bytes.Buffer
	for token := code.TtNothing; token <= code.TtQuote; token++ {
		escape.StyleText(&out, token.String(), styles.ResolveTokenStyle(token))
		out.WriteString("  ")
	}
	return out.String()
}

const visualMarkdown = "# Chernetka notes\n" + `
Plain text with **bold**, *italic*, ` + "`inline code`" + `, ~~struck~~ and a
backslash \*escape\* that stays literal.

Setext Heading
--------------

## Lists

- [x] markdown syntax highlight
- [ ] commit description highlight
  - nested item with *emphasis*
1. ordered item
2) also ordered

> A block quote with **emphasis** inside.
> ### and a heading nested in the quote

| Language | Backend      | Done |
|----------|--------------|:----:|
| Go       | tree-sitter  | yes  |
| Markdown | ` + "`mdScanner`" + `  | yes  |

` + "```go" + `
func main() {
	fmt.Println("inside a fence, not highlighted as Go")
}
` + "```" + `

    an indented code block

Links: [inline](https://example.com "Title"), [ref][r], <https://autolink.example>,
a bare https://go.dev/ref/spec URL, and an ![image](./demo.png).

[r]: https://example.com/reference "Reference Title"

---

An <b>HTML</b> tag, and snake_case_words that must not turn italic.
`

const visualGo = `// Package visual is here to be looked at.
package visual

import (
	"fmt"
	"strings"
)

// Greeter greets.
type Greeter struct {
	Name  string
	Count int
}

const defaultName = "world"

func (g *Greeter) Greet(w *strings.Builder) error {
	if g.Name == "" {
		g.Name = defaultName
	}
	for i := range g.Count {
		fmt.Fprintf(w, "hello %s\t#%d\n", g.Name, i+1)
	}
	return nil
}

var banner = ` + "`" + `a raw string
that spans
three lines` + "`" + `

/*
A block comment
across lines.
*/

func main() {
	g := &Greeter{Name: "chernetka", Count: 3}
	literals := []any{0x1f, 1.5e3, 'r', true, false, nil}
	_, _ = g, literals
}
`

const visualNix = `# A package, to be looked at.
{ lib, stdenv, fetchurl, enableTests ? true, ... }@args:

let
  version = "1.2.${toString 3}";
  greet = name: "hello \${name}\n";
  script = ''
    echo ${lib.escapeShellArg version} ''${literal}
  '';
in
stdenv.mkDerivation rec {
  pname = "chernetka";
  inherit version;
  src = fetchurl {
    url = "https://example.com/${pname}-${version}.tar.gz";
    hash = lib.fakeHash;
  };
  doCheck = enableTests && args ? lib;
  paths = [ ./src ~/config <nixpkgs> https://nixos.org ];
  meta = with lib; {
    license = licenses.mit or null;
    count = builtins.length [ 1 2.5 ];
    broken = if stdenv.isDarwin then false else null;
  };
}
`

const visualCUE = `// Package config is here to be looked at.
package config

import (
	"strings"
	"list"
)

// #Service describes a service.
#Service: {
	name!:     string & strings.MinRunes(1)
	port:      int | *8080
	replicas?: >=1 & <=10
	labels: [string]: string
	...
}

_#internal: {tags: [...string]}

svc: #Service & {
	name: "api"
	url:  "http://\(name):\(port + 1)/\n"
	if port > 80 {tls: true}
	sorted: list.Sort([3, 1.5], list.Ascending)
	count:  len(sorted)
	extra:  null | _ | _|_
	let base = "x"
	items: [for i, x in sorted {"\(base)-\(i)": x}]
	id: string @go(ID)
}
`

const visualJSON = `{
  // JSONC allows comments.
  "name": "chernetka",
  "version": 1.5e3,
  "escaped \"key\"": "tab\there \u00e9",
  "tags": ["editor", "go"],
  "nested": {"enabled": true, "disabled": false, "missing": null}
}
`

const visualYAML = `%YAML 1.2
---
# A workflow, to be looked at.
name: CI
on:
  push:
    branches: [main, "release/*"]
defaults: &defaults
  timeout: 30
  retries: 2.5
  enabled: yes
  nothing: ~
  since: 2026-09-26
jobs:
  test:
    <<: *defaults
    runs-on: 'ubuntu-latest'
    env: !!map {GOFLAGS: -mod=mod, "CGO": "1"}
    steps:
      - uses: actions/checkout@v4
      - run: |
          go test ./...
          echo "done\n"
...
`

const visualShell = `#!/usr/bin/env bash
# Build and test, to be looked at.
set -euo pipefail
export GOFLAGS=-mod=mod
readonly root="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

log() {
  local level=$1; shift
  printf '%s: %s\n' "$level" "$*" >&2
}

for pkg in ./cmd/... ./internal/...; do
  if [[ -d "$root/$pkg" && $# -gt 0 ]]; then
    go test -count=1 "$pkg" 2>/dev/null || log error "failed: $pkg ($?)"
  elif [ -z "${SKIP:-}" ]; then
    echo $'tab\there' $((2 + 3)) ` + "`date`" + `
  fi
done

case "$(uname)" in
  Darwin) open . ;;
  *) diff <(ls a) <(ls b) | grep -E '^[<>]' ;;
esac

cat <<EOF
Built in $root at $(date).
EOF
`

const visualC = `// Count the words of the files, to be looked at.
#include <stdio.h>
#include "words.h"

#define BUF_SIZE 4096
#define IS_SPACE(c) ((c) == ' ' || (c) == '\n')

typedef struct {
	const char *path;
	unsigned long words;
} result_t;

static int count(FILE *f, result_t *r) {
	char buf[BUF_SIZE];
	size_t n;
	int in_word = 0;
	while ((n = fread(buf, 1, sizeof buf, f)) > 0) {
		for (size_t i = 0; i < n; i++) {
			if (IS_SPACE(buf[i])) {
				in_word = 0;
			} else if (!in_word) {
				in_word = 1;
				r->words++;
			}
		}
	}
	return ferror(f) ? -1 : 0;
}

int main(int argc, char **argv) {
	for (int i = 1; i < argc; i++) {
		result_t r = {.path = argv[i], .words = 0};
		/* A missing file is reported, not fatal. */
		FILE *f = fopen(r.path, "r");
		if (f == NULL || count(f, &r) != 0) {
			fprintf(stderr, "%s: failed\n", r.path);
			continue;
		}
		printf("%8lu %s\t%c\n", r.words, r.path, 'w');
		fclose(f);
	}
	return 0;
}
`

const visualCPP = `// A small cache, to be looked at.
#pragma once
#include <memory>
#include <unordered_map>

namespace app::cache {

template <typename K, typename V>
class LRU final {
public:
	explicit LRU(std::size_t cap) noexcept : cap_(cap) {}
	~LRU() = default;

	[[nodiscard]] std::shared_ptr<V> get(const K& key) const;
	void put(K key, V value);
	bool operator==(const LRU& o) const = delete;

private:
	std::size_t cap_;
	std::unordered_map<K, std::shared_ptr<V>> items_;
};

template <typename K, typename V>
std::shared_ptr<V> LRU<K, V>::get(const K& key) const {
	if (auto it = items_.find(key); it != items_.end()) {
		return it->second;
	}
	return nullptr;
}

} // namespace app::cache

int main() {
	using namespace app::cache;
	auto c = std::make_unique<LRU<std::string, int>>(16);
	c->put("one", 1);
	constexpr auto msg = R"(raw "string")";
	try {
		std::cout << msg << u8"\u00e9" << 1.5e3f << '\n';
	} catch (const std::exception& e) {
		throw;
	}
	return this == nullptr ? 0 : 1;
}
`

const visualCSharp = `// A repository, to be looked at.
#nullable enable
using System;
using System.Collections.Generic;
using System.Linq;

namespace App.Data;

[Serializable]
public sealed record Item(int Id, string Name);

public interface IRepo<T> where T : class
{
    Task<T?> FindAsync(int id);
}

/// <summary>Keeps the items in memory.</summary>
public class MemoryRepo : IRepo<Item>, IDisposable
{
    private readonly List<Item> items = new();
    public int Count { get; private set; }
    public event EventHandler? Changed;

    public MemoryRepo(IEnumerable<Item> seed) => items.AddRange(seed);

    public async Task<Item?> FindAsync(int id)
    {
        await Task.Delay(1_000);
        var found = from i in items where i.Id == id select i;
        return found.FirstOrDefault() is { } item ? item : null;
    }

    public void Log(Item item)
    {
        Console.WriteLine($"{item.Id,5}: {item.Name} \t{DateTime.Now:HH:mm}");
        Console.WriteLine(@"C:\data" + """raw "text" """ + 'x');
    }

    public void Dispose() { }

    enum State { Open = 0x1, Closed = 2 }
}
`

const visualGitMsg = `some/path: change title

Test commit message

Signed-off-by: chernetka@rmazur.io
Another-tag: something

P.S.

# Some comment 1
# Some comment 2
`

const visualGitRebase = `
p abcdef # pick a commit
edit ab123 # edit a commit
x go test ./... # verify
f another # continue
squash one more
# this should be it
something irrelevant
`
