package extsyntaxhl

import "testing"

func TestDockerfileHighlight(t *testing.T) {
	doc := hlDoc{
		// The parser directives at the top read as directives, unlike the comments after them.
		{"# syntax=docker/dockerfile:1", []string{"0:28:Directive"}},
		{"# check=error=true", []string{"0:18:Directive"}},
		{"", nil},
		{"# The build stage.", []string{"0:18:Comment"}},
		{"ARG GO_VERSION=1.23", []string{"0:3:Keyword", "4:14:Field"}},
		// The image of a stage reads as an import and its name as a declaration, the flags as constants.
		{"FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build", []string{"0:4:Keyword", "5:15:Constant", "16:17:Escape", "17:30:Field", "31:38:ImportRef", "38:40:Escape", "40:50:Field", "50:51:Escape", "52:54:Keyword", "55:60:FuncDeclaration"}},
		{"WORKDIR /src", []string{"0:7:Keyword"}},
		{"RUN --mount=type=cache,target=/go/pkg/mod \\", []string{"0:3:Keyword", "4:11:Constant"}},
		{"    --mount=type=bind,source=go.mod,target=go.mod \\", []string{"4:11:Constant"}},
		{"    go mod download -x", nil},
		// The body of a heredoc is a string, its variables expanded unless its word is quoted.
		{"COPY <<EOF /etc/app.conf", []string{"0:4:Keyword", "5:10:StringLiteral"}},
		{"port=${PORT:-8080}", []string{"0:5:StringLiteral", "5:7:Escape", "7:11:Field", "11:17:StringLiteral", "17:18:Escape"}},
		{"EOF", []string{"0:3:StringLiteral"}},
		{"RUN <<-'SH' sh", []string{"0:3:Keyword", "4:11:StringLiteral"}},
		{"\techo \"$HOME\"", []string{"0:13:StringLiteral"}},
		{"\tSH", []string{"0:3:StringLiteral"}},
		{"", nil},
		{"FROM gcr.io/distroless/static AS final", []string{"0:4:Keyword", "5:29:ImportRef", "30:32:Keyword", "33:38:FuncDeclaration"}},
		// The names of LABEL and ENV read as properties, wherever their lines are continued.
		{"LABEL org.opencontainers.image.source=\"https://example.com\" version=1", []string{"0:5:Keyword", "6:37:Field", "38:59:StringLiteral", "60:67:Field"}},
		{"ENV TZ=UTC \\", []string{"0:3:Keyword", "4:6:Field"}},
		{"    PATH=\"/app:$PATH\"", []string{"4:8:Field", "9:15:StringLiteral", "15:16:Escape", "16:20:Field", "20:21:StringLiteral"}},
		{"COPY --from=build --chown=nonroot /out/app /app", []string{"0:4:Keyword", "5:11:Constant", "18:25:Constant"}},
		{"EXPOSE 8080/tcp", []string{"0:6:Keyword", "7:11:NumberLiteral"}},
		{"USER nonroot", []string{"0:4:Keyword"}},
		// The exec form is a JSON array of strings.
		{"HEALTHCHECK --interval=30s CMD [\"/app\", \"health\"]", []string{"0:11:Keyword", "12:22:Constant", "27:30:Keyword", "32:38:StringLiteral", "40:48:StringLiteral"}},
		{"ENTRYPOINT [\"/app\"]", []string{"0:10:Keyword", "12:18:StringLiteral"}},
	}

	buf, ext := openDoc(t, "Dockerfile", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}
