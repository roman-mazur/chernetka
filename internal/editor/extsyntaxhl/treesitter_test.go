package extsyntaxhl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unicode"

	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
)

func TestGoHighlight(t *testing.T) {
	doc := hlDoc{
		{"// Command comment.", []string{"0:19:Comment"}},
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"import \"fmt\"", []string{"0:6:Keyword", "7:12:ImportRef"}},
		{"", nil},
		{"type T struct{ N int }", []string{
			"0:4:Keyword", "5:6:TypeRef", "7:13:Keyword", "15:16:Field", "17:20:TypeRef",
		}},
		{"", nil},
		{"func (t *T) Do(s string) error {", []string{
			"0:4:Keyword", "6:7:Identifier", "9:10:TypeRef",
			"12:14:FuncDeclaration", "15:16:Identifier", "17:23:TypeRef", "25:30:TypeRef",
		}},
		{"\tconst x = 1.5", []string{"1:6:Keyword", "7:8:Identifier", "11:14:NumberLiteral"}},
		// The \t escape sequence sits inside the string literal and overrides it.
		{"\tfmt.Println(\"a\\tb\", t.N, x, nil)", []string{
			"1:4:Identifier", "5:12:Call", "13:15:StringLiteral", "15:17:Escape",
			"17:19:StringLiteral", "21:22:Identifier", "23:24:Field", "26:27:Identifier",
			"29:32:Constant",
		}},
		{"\treturn nil", []string{"1:7:Keyword", "8:11:Constant"}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestGoMultiLineLiteral covers a node that starts on one line and ends on
// another: it has to be cut into one span per line, each clipped to its own
// line, instead of being reported with the start line's number and the end
// line's column.
func TestGoMultiLineLiteral(t *testing.T) {
	doc := hlDoc{
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"var raw = `first", []string{"0:3:Keyword", "4:7:Identifier", "10:16:StringLiteral"}},
		{"second line is longer", []string{"0:21:StringLiteral"}},
		{"third`", []string{"0:6:StringLiteral"}},
		{"", nil},
		{"/* a block", []string{"0:10:Comment"}},
		{"   comment */", []string{"0:13:Comment"}},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestGoHighlightAfterEdit(t *testing.T) {
	doc := hlDoc{
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"func main() {}", []string{"0:4:Keyword", "5:9:FuncDeclaration"}},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	hl := highlighterOf(t, buf, ext)
	doc.check(t, hl)

	// Inserting a line shifts everything below it down by one.
	insertLine(t, buf, ext, 1)
	shifted := hlDoc{doc[0], {"", nil}, doc[1], doc[2]}
	shifted.check(t, hl)
}

func TestNixHighlight(t *testing.T) {
	doc := hlDoc{
		{"# A comment.", []string{"0:12:Comment"}},
		{"{ lib, enable ? true, ... }@args:", []string{
			"2:5:Identifier", "7:13:Identifier", "16:20:Constant", "28:32:Identifier",
		}},
		{"let", []string{"0:3:Keyword"}},
		{"  greet = name: \"hi ${name}\\n\";", []string{
			"2:7:FuncDeclaration", "10:14:Identifier", "16:20:StringLiteral",
			"20:22:Escape", "22:26:Identifier", "26:29:Escape", "29:30:StringLiteral",
		}},
		{"in rec {", []string{"0:2:Keyword", "3:6:Keyword"}},
		{"  a.b = lib.mkIf enable [ ./src 1.5 ];", []string{
			"2:3:Field", "4:5:Field", "8:11:Identifier", "12:16:Call", "17:23:Identifier",
			"26:31:StringLiteral", "32:35:NumberLiteral",
		}},
		{"  inherit (builtins) toString;", []string{"2:9:Keyword", "11:19:ImportRef", "21:29:Field"}},
		{"  c = args.x or null;", []string{
			"2:3:Field", "6:10:Identifier", "11:12:Field", "13:15:Keyword", "16:20:Constant",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "default.nix", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestCUEHighlight(t *testing.T) {
	doc := hlDoc{
		{"package config", []string{"0:7:Keyword", "8:14:Identifier"}},
		{"", nil},
		{"import \"strings\"", []string{"0:6:Keyword", "7:16:ImportRef"}},
		{"", nil},
		{"// #Service is a definition.", []string{"0:28:Comment"}},
		{"#Service: {", []string{"0:8:TypeRef"}},
		{"\tname!:  string & strings.MinRunes(1)", []string{
			"1:5:Field", "9:15:TypeRef", "18:25:Identifier", "26:34:Call", "35:36:NumberLiteral",
		}},
		{"\tport?:  int | *8080 @go(Port)", []string{
			"1:5:Field", "9:12:TypeRef", "16:20:NumberLiteral", "21:30:Annotation",
		}},
		{"}", nil},
		{"", nil},
		{"svc: #Service & {", []string{"0:3:Field", "5:13:TypeRef"}},
		// The interpolated expression is code, so the string doesn't color it.
		{"\turl: \"http://\\(name):\\(port + 1)\\n\"", []string{
			"1:4:Field", "6:14:StringLiteral", "14:16:Escape", "16:20:Identifier", "20:21:Escape",
			"21:22:StringLiteral", "22:24:Escape", "24:31:Identifier", "31:32:NumberLiteral",
			"32:35:Escape", "35:36:StringLiteral",
		}},
		{"\tif port > 80 {tls: true}", []string{
			"1:3:Keyword", "4:8:Identifier", "11:13:NumberLiteral", "15:18:Field", "20:24:Constant",
		}},
		{"\tn: len([for x in [1, 2.5] {x}])", []string{
			"1:2:Field", "4:7:Call", "9:12:Keyword", "13:14:Identifier", "15:17:Keyword",
			"19:20:NumberLiteral", "22:25:NumberLiteral", "28:29:Identifier",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "config.cue", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestCUEErrorRecovery checks the input the core v0.25.0 never finished parsing
// with the CUE grammar.
func TestCUEErrorRecovery(t *testing.T) {
	const text = "{R(z&["
	done := make(chan struct{})
	go func() {
		defer close(done)
		openDoc(t, "recovery.cue", text)
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("parsing does not finish")
	}
}

func TestJSONHighlight(t *testing.T) {
	doc := hlDoc{
		{"{", nil},
		{"  // A comment.", []string{"2:15:Comment"}},
		{"  \"a\\\"b\": \"x\\ty\",", []string{
			"2:4:Field", "4:6:Escape", "6:8:Field", "10:12:StringLiteral", "12:14:Escape", "14:16:StringLiteral",
		}},
		{"  \"n\": [1.5e3, true, false, null]", []string{
			"2:5:Field", "8:13:NumberLiteral", "15:19:Constant", "21:26:Constant", "28:32:Constant",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "package.json", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestYAMLHighlight(t *testing.T) {
	doc := hlDoc{
		{"---", []string{"0:3:Keyword"}},
		{"# A comment.", []string{"0:12:Comment"}},
		{"name: CI", []string{"0:4:Field", "6:8:StringLiteral"}},
		{"base: &base", []string{"0:4:Field", "6:11:Constant"}},
		{"  n: 30", []string{"2:3:Field", "5:7:NumberLiteral"}},
		{"  f: 2.5", []string{"2:3:Field", "5:8:NumberLiteral"}},
		{"  ok: true", []string{"2:4:Field", "6:10:Constant"}},
		{"  \"quoted key\": 'v'", []string{"2:14:Field", "16:19:StringLiteral"}},
		{"job:", []string{"0:3:Field"}},
		{"  <<: *base", []string{"2:4:Field", "6:11:Constant"}},
		{"  env: !!map {A: \"1\\n\"}", []string{
			"2:5:Field", "7:12:TypeRef", "14:15:Field", "17:19:StringLiteral", "19:21:Escape", "21:22:StringLiteral",
		}},
		{"  steps: [checkout, test]", []string{"2:7:Field", "10:18:StringLiteral", "20:24:StringLiteral"}},
	}

	buf, ext := openDoc(t, "ci.yaml", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestShellHighlight(t *testing.T) {
	doc := hlDoc{
		{"#!/bin/sh", []string{"0:9:Directive"}},
		{"export DIR=\"${1:-.}\"", []string{
			"0:6:Keyword", "7:10:Field", "11:12:StringLiteral", "12:14:Escape", "14:15:Constant",
			"15:18:Identifier", "18:19:Escape", "19:20:StringLiteral",
		}},
		{"run() { ls -la \"$DIR\" >&2; }", []string{
			"0:3:FuncDeclaration", "8:10:Call", "11:14:Constant",
			"15:16:StringLiteral", "16:17:Escape", "17:20:Field", "20:21:StringLiteral", "24:25:NumberLiteral",
		}},
		// A substitution in a string is code again, not string content.
		{"if [ -n \"x $(date)\" ]; then", []string{
			"0:2:Keyword", "5:7:Keyword", "8:11:StringLiteral", "11:13:Escape", "13:17:Call",
			"17:18:Escape", "18:19:StringLiteral", "23:27:Keyword",
		}},
		{"  echo 'a $b' $# 42", []string{
			"2:6:Call", "7:13:StringLiteral", "14:15:Escape", "15:16:Constant", "17:19:NumberLiteral",
		}},
		{"fi", []string{"0:2:Keyword"}},
	}

	buf, ext := openDoc(t, "run.sh", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestCHighlight(t *testing.T) {
	doc := hlDoc{
		{"#include <stdio.h>", []string{"0:8:Directive", "9:18:ImportRef"}},
		{"#define MAX_LEN 16", []string{"0:7:Directive", "8:15:Constant"}},
		{"", nil},
		{"typedef struct { int x; unsigned long y; } point_t;", []string{
			"0:7:Keyword", "8:14:Keyword", "17:20:TypeRef", "21:22:Field", "24:37:TypeRef", "38:39:Field", "43:50:TypeRef",
		}},
		{"", nil},
		{"static const char *name(point_t *p) {", []string{
			"0:6:Keyword", "7:12:Keyword", "13:17:TypeRef", "19:23:FuncDeclaration", "24:31:TypeRef", "33:34:Identifier",
		}},
		// Upper case names read as macros.
		{"\tif (p->x > MAX_LEN) return NULL;", []string{
			"1:3:Keyword", "5:6:Identifier", "8:9:Field", "12:19:Constant", "21:27:Keyword", "28:32:Constant",
		}},
		{"\tprintf(\"%d\\n\", p->x, 'a');", []string{
			"1:7:Call", "8:11:StringLiteral", "11:13:Escape", "13:14:StringLiteral", "16:17:Identifier", "19:20:Field",
			"22:25:StringLiteral",
		}},
		{"\tp->fn(sizeof(int), 0x1fu); // done", []string{
			"1:2:Identifier", "4:6:Call", "7:13:Keyword", "14:17:TypeRef", "20:25:NumberLiteral", "28:35:Comment",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "main.c", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestCPPHighlight(t *testing.T) {
	doc := hlDoc{
		{"namespace app {", []string{"0:9:Keyword", "10:13:ImportRef"}},
		{"template <typename T>", []string{"0:8:Keyword", "10:18:Keyword", "19:20:TypeRef"}},
		{"class Box final : public Base {", []string{
			"0:5:Keyword", "6:9:TypeRef", "10:15:Keyword", "18:24:Keyword", "25:29:TypeRef",
		}},
		{"\texplicit Box(T v) noexcept : v_(v) {}", []string{
			"1:9:Keyword", "10:13:FuncDeclaration", "14:15:TypeRef", "16:17:Identifier", "19:27:Keyword", "30:32:Field",
			"33:34:Identifier",
		}},
		{"\t~Box() override = default;", []string{"2:5:FuncDeclaration", "8:16:Keyword", "19:26:Keyword"}},
		{"\tauto get() const -> const T& { return v_; }", []string{
			"1:5:TypeRef", "6:9:FuncDeclaration", "12:17:Keyword", "21:26:Keyword", "27:28:TypeRef", "32:38:Keyword",
			"39:41:Identifier",
		}},
		// An attribute reads apart from the code, brackets included.
		{"\t[[nodiscard, gnu::pure]] int size() const;", []string{
			"1:12:Annotation", "14:17:Annotation", "19:25:Annotation", "26:29:TypeRef", "30:34:FuncDeclaration", "37:42:Keyword",
		}},
		{"};", nil},
		{"}", nil},
		{"void app::run() {", []string{"0:4:TypeRef", "5:8:ImportRef", "10:13:FuncDeclaration"}},
		{"\tauto b = std::make_unique<Box<int>>(1);", []string{
			"1:5:TypeRef", "6:7:Identifier", "10:13:ImportRef", "15:26:Call", "27:30:TypeRef", "31:34:TypeRef",
			"37:38:NumberLiteral",
		}},
		{"\tstd::cout << R\"(raw)\" << nullptr << this->x;", []string{
			"1:4:ImportRef", "6:10:Identifier", "14:22:StringLiteral", "26:33:Constant", "37:41:Keyword", "43:44:Field",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "box.cpp", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestCSharpHighlight(t *testing.T) {
	doc := hlDoc{
		// A dotted namespace is one span, not a chain of identifiers.
		{"using System.Collections.Generic;", []string{"0:5:Keyword", "6:32:ImportRef"}},
		{"", nil},
		{"[Serializable]", []string{"0:14:Annotation"}},
		{"public class Repo<T> : IRepo<T> where T : class", []string{
			"0:6:Keyword", "7:12:Keyword", "13:17:TypeRef", "18:19:TypeRef", "23:28:TypeRef", "29:30:TypeRef",
			"32:37:Keyword", "42:47:Keyword",
		}},
		{"{", nil},
		{"    public int Count { get; private set; }", []string{
			"4:10:Keyword", "11:14:TypeRef", "15:20:Field", "23:26:Keyword", "28:35:Keyword", "36:39:Keyword",
		}},
		{"    public async Task<T?> FindAsync(int id)", []string{
			"4:10:Keyword", "11:16:Keyword", "17:21:TypeRef", "22:23:TypeRef", "26:35:FuncDeclaration", "36:39:TypeRef",
		}},
		{"    {", nil},
		{"        var x = items.Find(i => i is null);", []string{
			"8:11:Keyword", "22:26:Call", "34:36:Keyword", "37:41:Constant",
		}},
		// An interpolated expression is code again, not string content.
		{"        Console.WriteLine($\"id={id} \\t{x}\" + 'c');", []string{
			"16:25:Call", "26:31:StringLiteral", "31:32:Escape", "32:34:Identifier", "34:35:Escape", "35:36:StringLiteral",
			"36:39:Escape", "39:40:Identifier", "40:41:Escape", "41:42:StringLiteral", "45:48:StringLiteral",
		}},
		{"        return default;", []string{"8:14:Keyword", "15:22:Keyword"}},
		{"    }", nil},
		{"    enum Color { Red, Green }", []string{"4:8:Keyword", "9:14:TypeRef", "17:20:Constant", "22:27:Constant"}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "Repo.cs", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestJavaHighlight(t *testing.T) {
	doc := hlDoc{
		{"package app.util;", []string{"0:7:Keyword", "8:16:ImportRef"}},
		// A dotted package is one span, not a chain of identifiers.
		{"import java.util.List;", []string{"0:6:Keyword", "7:21:ImportRef"}},
		{"", nil},
		{"@Service", []string{"0:8:Annotation"}},
		{"public final class Repo<T> extends Base implements Store {", []string{"0:6:Keyword", "7:12:Keyword", "13:18:Keyword", "19:23:TypeRef", "24:25:TypeRef", "27:34:Keyword", "35:39:TypeRef", "40:50:Keyword", "51:56:TypeRef"}},
		{"    private static final int MAX_SIZE = 0x10;", []string{"4:11:Keyword", "12:18:Keyword", "19:24:Keyword", "25:28:TypeRef", "29:37:Constant", "40:44:NumberLiteral"}},
		{"    public Repo(String name) { this.name = name; }", []string{"4:10:Keyword", "11:15:FuncDeclaration", "16:22:TypeRef", "31:35:Keyword", "36:40:Field"}},
		{"    @Override", []string{"4:13:Annotation"}},
		{"    public List<T> find(int id) throws IOException {", []string{"4:10:Keyword", "11:15:TypeRef", "16:17:TypeRef", "19:23:FuncDeclaration", "24:27:TypeRef", "32:38:Keyword", "39:50:TypeRef"}},
		{"        var s = \"id\\t\" + id + 'c' + 1.5f;", []string{"8:11:TypeRef", "16:19:StringLiteral", "19:21:Escape", "21:22:StringLiteral", "30:33:StringLiteral", "36:40:NumberLiteral"}},
		// A capitalized name a member is taken from reads as a type, an upper case one as a constant.
		{"        System.out.println(Math.max(id, Color.RED.ordinal()));", []string{"8:14:TypeRef", "15:18:Field", "19:26:Call", "27:31:TypeRef", "32:35:Call", "40:45:TypeRef", "46:49:Constant", "50:57:Call"}},
		{"        return items.stream().filter(x -> x != null).toList(); // done", []string{"8:14:Keyword", "21:27:Call", "30:36:Call", "47:51:Constant", "53:59:Call", "63:70:Comment"}},
		{"    }", nil},
		{"    enum Color { RED, GREEN }", []string{"4:8:Keyword", "9:14:TypeRef", "17:20:Constant", "22:27:Constant"}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "Repo.java", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestKotlinHighlight(t *testing.T) {
	doc := hlDoc{
		{"package app.util", []string{"0:7:Keyword", "8:11:ImportRef", "12:16:ImportRef"}},
		// The parts of a dotted name are colored, not the dots.
		{"import kotlin.math.max", []string{"0:6:Keyword", "7:13:ImportRef", "14:18:ImportRef", "19:22:ImportRef"}},
		{"", nil},
		{"@JvmInline", []string{"0:10:Annotation"}},
		{"data class Point(val x: Int, var y: Int = 0) {", []string{"0:4:Keyword", "5:10:Keyword", "11:16:TypeRef", "17:20:Keyword", "21:22:Field", "24:27:TypeRef", "29:32:Keyword", "33:34:Field", "36:39:TypeRef", "42:43:NumberLiteral"}},
		// A string template is code again, not string content.
		{"    private val label = \"p($x, ${x + y})\\n\"", []string{"4:11:Keyword", "12:15:Keyword", "16:21:Field", "24:27:StringLiteral", "27:28:Escape", "28:29:Identifier", "29:31:StringLiteral", "31:33:Escape", "33:38:Identifier", "38:39:Escape", "39:43:StringLiteral"}},
		{"    fun String.shout(n: Int): String? {", []string{"4:7:Keyword", "8:14:TypeRef", "15:20:FuncDeclaration", "21:22:Identifier", "24:27:TypeRef", "30:36:TypeRef"}},
		{"        if (n !in 0..10) return null", []string{"8:10:Keyword", "12:13:Identifier", "14:17:Keyword", "18:19:NumberLiteral", "21:23:NumberLiteral", "25:31:Keyword", "32:36:Constant"}},
		{"        val list = listOf(1, 2L, 0xFFu)", []string{"8:11:Keyword", "12:16:Identifier", "19:25:Call", "26:27:NumberLiteral", "29:31:NumberLiteral", "33:38:NumberLiteral"}},
		{"        println(list.map { it * 2 }.sum())", []string{"8:15:Call", "16:20:Identifier", "21:24:Call", "27:29:Identifier", "32:33:NumberLiteral", "36:39:Call"}},
		{"        return Color.RED.name + 'c' // done", []string{"8:14:Keyword", "15:20:TypeRef", "21:24:Constant", "25:29:Field", "32:35:StringLiteral", "36:43:Comment"}},
		{"    }", nil},
		{"}", nil},
		{"enum class Color { RED, GREEN }", []string{"0:4:Keyword", "5:10:Keyword", "11:16:TypeRef", "19:22:Constant", "24:29:Constant"}},
	}

	buf, ext := openDoc(t, "Point.kt", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestTypeScriptHighlight(t *testing.T) {
	doc := hlDoc{
		{"import { readFile } from \"node:fs/promises\";", []string{"0:6:Keyword", "9:17:Identifier", "20:24:Keyword", "25:43:ImportRef"}},
		{"", nil},
		// A decorator reads as an annotation, with its @ but not its arguments.
		{"@Injectable({ root: true })", []string{"0:11:Annotation", "14:18:Field", "20:24:Constant"}},
		{"export class Repo<T extends Item> implements Store {", []string{"0:6:Keyword", "7:12:Keyword", "13:17:TypeRef", "18:19:TypeRef", "20:27:Keyword", "28:32:TypeRef", "34:44:Keyword", "45:50:TypeRef"}},
		{"  private static readonly MAX_SIZE = 0x10;", []string{"2:9:Keyword", "10:16:Keyword", "17:25:Keyword", "26:34:Constant", "37:41:NumberLiteral"}},
		{"  constructor(private name: string) {}", []string{"2:13:FuncDeclaration", "14:21:Keyword", "22:26:Identifier", "28:34:TypeRef"}},
		{"  async find(id: number): Promise<T | undefined> {", []string{"2:7:Keyword", "8:12:FuncDeclaration", "13:15:Identifier", "17:23:TypeRef", "26:33:TypeRef", "34:35:TypeRef", "38:47:Constant"}},
		// A template substitution is code again, not string content.
		{"    const s = `id\\t${id + 1}`;", []string{"4:9:Keyword", "10:11:Identifier", "14:17:StringLiteral", "17:21:Escape", "21:26:Identifier", "26:27:NumberLiteral", "27:28:Escape", "28:29:StringLiteral"}},
		// A capitalized name reads as a type, an upper case one as a constant.
		{"    return this.items.find((x) => x.id === Repo.MAX_SIZE) ?? null; // done", []string{"4:10:Keyword", "11:15:Keyword", "16:21:Field", "22:26:Call", "28:29:Identifier", "34:35:Identifier", "36:38:Field", "43:47:TypeRef", "48:56:Constant", "61:65:Constant", "67:74:Comment"}},
		{"  }", nil},
		{"}", nil},
		{"enum Color { Red = \"red\", Green }", []string{"0:4:Keyword", "5:10:TypeRef", "13:16:Constant", "19:24:StringLiteral", "26:31:Constant"}},
		{"type Pair = [key: keyof Item, value?: unknown];", []string{"0:4:Keyword", "5:9:TypeRef", "13:16:Identifier", "18:23:Keyword", "24:28:TypeRef", "30:35:Identifier", "38:45:TypeRef"}},
	}

	buf, ext := openDoc(t, "repo.ts", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestTSXHighlight(t *testing.T) {
	doc := hlDoc{
		// A directive prologue reads as a directive.
		{"\"use client\";", []string{"0:12:Directive"}},
		{"export const App = ({ title }: Props) => (", []string{"0:6:Keyword", "7:12:Keyword", "13:16:FuncDeclaration", "31:36:TypeRef"}},
		// A platform element reads like a keyword, a component as a type.
		{"  <div className=\"app\" hidden>", []string{"3:6:Keyword", "7:16:Field", "17:22:StringLiteral", "23:29:Field"}},
		{"    <Header title={title.toUpperCase()} />", []string{"5:11:TypeRef", "12:17:Field", "19:24:Identifier", "25:36:Call"}},
		// The text of an element is left alone, but for its character references.
		{"    Tom &amp; Jerry", []string{"8:13:Escape"}},
		{"  </div>", []string{"4:7:Keyword"}},
		{");", nil},
	}

	buf, ext := openDoc(t, "App.tsx", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestJavaScriptHighlight(t *testing.T) {
	doc := hlDoc{
		{"#!/usr/bin/env node", []string{"0:19:Directive"}},
		{"const { join } = require(\"path\");", []string{"0:5:Keyword", "17:24:Call", "25:31:StringLiteral"}},
		{"function main(argv) {", []string{"0:8:Keyword", "9:13:FuncDeclaration", "14:18:Identifier"}},
		// JavaScript is read with the TSX grammar, JSX included.
		{"  return /a+b/g.test(argv[2]) ? <b>yes</b> : 42n;", []string{"2:8:Keyword", "9:15:StringLiteral", "16:20:Call", "21:25:Identifier", "26:27:NumberLiteral", "33:34:Keyword", "40:41:Keyword", "45:48:NumberLiteral"}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "main.js", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestRustHighlight(t *testing.T) {
	doc := hlDoc{
		// An attribute reads as an annotation, with its #, ! and brackets but not its arguments.
		{"#![allow(dead_code)]", []string{"0:8:Annotation", "9:18:Identifier", "19:20:Annotation"}},
		{"use std::collections::HashMap;", []string{"0:3:Keyword", "4:7:ImportRef", "9:20:ImportRef", "22:29:TypeRef"}},
		// The segments of a path before its last one are modules.
		{"use crate::util::{self, Config};", []string{"0:3:Keyword", "4:9:Keyword", "11:15:ImportRef", "18:22:Keyword", "24:30:TypeRef"}},
		{"", nil},
		{"const MAX_SIZE: usize = 0x10;", []string{"0:5:Keyword", "6:14:Constant", "16:21:TypeRef", "24:28:NumberLiteral"}},
		{"", nil},
		{"#[derive(Debug, Clone)]", []string{"0:8:Annotation", "9:14:TypeRef", "16:21:TypeRef", "22:23:Annotation"}},
		// A lifetime reads as a type, its quote included.
		{"pub struct Repo<'a, T> {", []string{"0:3:Keyword", "4:10:Keyword", "11:15:TypeRef", "16:18:TypeRef", "20:21:TypeRef"}},
		{"    name: &'a str,", []string{"4:8:Field", "11:13:TypeRef", "14:17:TypeRef"}},
		{"    items: Vec<T>,", []string{"4:9:Field", "11:14:TypeRef", "15:16:TypeRef"}},
		{"}", nil},
		{"", nil},
		{"impl<'a, T: Clone> Repo<'a, T> {", []string{"0:4:Keyword", "5:7:TypeRef", "9:10:TypeRef", "12:17:TypeRef", "19:23:TypeRef", "24:26:TypeRef", "28:29:TypeRef"}},
		{"    pub fn find(&self, id: u32) -> Option<&T> {", []string{"4:7:Keyword", "8:10:Keyword", "11:15:FuncDeclaration", "17:21:Keyword", "23:25:Identifier", "27:30:TypeRef", "35:41:TypeRef", "43:44:TypeRef"}},
		{"        let s = format!(\"id\\t{}\", id + 1);", []string{"8:11:Keyword", "12:13:Identifier", "16:23:Call", "24:27:StringLiteral", "27:29:Escape", "29:32:StringLiteral", "34:36:Identifier", "39:40:NumberLiteral"}},
		// A macro invocation reads as a call, its ! included.
		{"        println!(\"{s} {}\", MAX_SIZE);", []string{"8:16:Call", "17:25:StringLiteral", "27:35:Constant"}},
		{"        self.items.iter().find(|x| x.id == id).or(None) // done", []string{"8:12:Keyword", "13:18:Field", "19:23:Call", "26:30:Call", "32:33:Identifier", "35:36:Identifier", "37:39:Field", "43:45:Identifier", "47:49:Call", "50:54:TypeRef", "56:63:Comment"}},
		{"    }", nil},
		{"}", nil},
		{"", nil},
		{"#[tokio::main]", []string{"0:7:Annotation", "9:14:Annotation"}},
		{"async fn main() -> Result<(), Box<dyn std::error::Error>> {", []string{"0:5:Keyword", "6:8:Keyword", "9:13:FuncDeclaration", "19:25:TypeRef", "30:33:TypeRef", "34:37:Keyword", "38:41:ImportRef", "43:48:ImportRef", "50:55:TypeRef"}},
		{"    let mut m = HashMap::<String, i64>::new();", []string{"4:7:Keyword", "8:11:Keyword", "12:13:Identifier", "16:23:TypeRef", "26:32:TypeRef", "34:37:TypeRef", "40:43:Call"}},
		{"    match std::cmp::Ord::cmp(&1, &2) { std::cmp::Ordering::Less => true, _ => false };", []string{"4:9:Keyword", "10:13:ImportRef", "15:18:ImportRef", "20:23:TypeRef", "25:28:Call", "30:31:NumberLiteral", "34:35:NumberLiteral", "39:42:ImportRef", "44:47:ImportRef", "49:57:TypeRef", "59:63:TypeRef", "67:71:Constant", "78:83:Constant"}},
		// A capitalized call is an enum variant or a tuple struct, which reads as a type.
		{"    Ok(())", []string{"4:6:TypeRef"}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "main.rs", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestTreeSitterQueryHighlight(t *testing.T) {
	doc := hlDoc{
		{"; Calls.", []string{"0:8:Comment"}},
		// A node type reads as a type, a field as a field, a capture as a constant.
		{"(call_expression", []string{"1:16:TypeRef"}},
		{"  function: (identifier) @function.call)", []string{"2:10:Field", "13:23:TypeRef", "25:39:Constant"}},
		{"", nil},
		{"((identifier) @constant", []string{"2:12:TypeRef", "14:23:Constant"}},
		// A predicate reads as a directive.
		{"  (#match? @constant \"^[A-Z]\\\\d+$\"))", []string{
			"3:10:Directive", "11:20:Constant", "21:28:StringLiteral", "28:30:Escape", "30:34:StringLiteral",
		}},
		{"", nil},
		{"[", nil},
		{"  \"return\"", []string{"2:10:StringLiteral"}},
		{"  (true)", []string{"3:7:TypeRef"}},
		{"] @keyword", []string{"2:10:Constant"}},
		{"", nil},
		// Anchors, quantifiers and wildcards change what a pattern matches.
		{"(block . (comment)* @doc !type _)", []string{
			"1:6:TypeRef", "7:8:Keyword", "10:17:TypeRef", "18:19:Keyword", "20:24:Constant", "25:30:Field", "31:32:Keyword",
		}},
		{"(#set! priority 90)", []string{"1:6:Directive"}},
		{"(MISSING identifier) @error", []string{"1:8:Keyword", "9:19:TypeRef", "21:27:Constant"}},
	}

	buf, ext := openDoc(t, "queries/go.scm", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestQueriesParse reads the highlight queries of this package with the query
// grammar, which they must parse with without an error.
func TestQueriesParse(t *testing.T) {
	entries, err := os.ReadDir("queries")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		path := filepath.Join("queries", entry.Name())
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		buf, ext := openDoc(t, path, string(src))
		doc, ok := buf.ExtensionData(ext.ID()).(*document)
		if !ok {
			t.Fatalf("%s is not highlighted", path)
		}
		doc.ensureParsed()
		if tree := doc.hl.(*tsHighlighter).tree; tree == nil || tree.RootNode().HasError() {
			t.Errorf("%s doesn't parse as a tree-sitter query", path)
		}
	}
}

func TestCommonLispHighlight(t *testing.T) {
	doc := hlDoc{
		{";;; Words, to be looked at.", []string{"0:27:Comment"}},
		// The names of the definitions read by what they define.
		{"(in-package :words)", []string{"1:11:Keyword", "12:18:ImportRef"}},
		{"", nil},
		{"(defparameter *top* 10)", []string{"1:13:Keyword", "14:19:Constant", "20:22:NumberLiteral"}},
		{"(defconstant +tab+ #\\Tab)", []string{"1:12:Keyword", "13:18:Constant", "19:24:StringLiteral"}},
		{"", nil},
		// The first element of a lambda list or a binding list is a name, not a call.
		{"(defun count-words (path &optional (limit *top*))", []string{"1:6:Keyword", "7:18:FuncDeclaration", "20:24:Identifier", "25:34:Keyword", "36:41:Identifier", "42:47:Constant"}},
		{"  \"Counts the words of the file at PATH.\"", []string{"2:41:StringLiteral"}},
		{"  (let ((table (make-hash-table :test #'equal)))", []string{"3:6:Keyword", "9:14:Identifier", "16:31:Call", "32:37:Constant", "40:45:Call"}},
		{"    (with-open-file (in path)", []string{"5:19:Keyword", "21:23:Identifier", "24:28:Identifier"}},
		{"      (loop for line = (read-line in nil)", []string{"7:11:Keyword", "12:15:Keyword", "16:20:Identifier", "21:22:Keyword", "24:33:Call", "34:36:Identifier", "37:40:Constant"}},
		{"            while line", []string{"12:17:Keyword", "18:22:Identifier"}},
		{"            do (incf (gethash line table 0))))", []string{"12:14:Keyword", "16:20:Keyword", "22:29:Call", "30:34:Identifier", "35:40:Identifier", "41:42:NumberLiteral"}},
		// A format directive reads like an escape in its string.
		{"    (format t \"~A words~%\" (hash-table-count table))", []string{"5:11:Call", "12:13:Constant", "14:15:StringLiteral", "15:17:Escape", "17:23:StringLiteral", "23:25:Escape", "25:26:StringLiteral", "28:44:Call", "45:50:Identifier"}},
		// A quoted list is data: its first element is no call.
		{"    (cl:subseq '(a b c) 0 limit)))", []string{"5:7:ImportRef", "8:14:Call", "17:18:Identifier", "19:20:Identifier", "21:22:Identifier", "24:25:NumberLiteral", "26:31:Identifier"}},
		{"", nil},
		// A reader conditional reads as a directive.
		{"#+sbcl (sb-ext:gc :full t)", []string{"0:6:Directive", "8:14:ImportRef", "15:17:Call", "18:23:Constant", "24:25:Constant"}},
		// The grammar takes only the lower case definitions, the query any case.
		{"(DEFUN Shout (x) (string-upcase x))", []string{"1:6:Keyword", "7:12:FuncDeclaration", "14:15:Identifier", "18:31:Call", "32:33:Identifier"}},
	}

	buf, ext := openDoc(t, "words.lisp", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestDirectives covers the instructions to the toolchain, which read apart
// from both the comments and the keywords, and the look-alikes that do not.
func TestDirectives(t *testing.T) {
	for path, doc := range map[string]hlDoc{
		"main.go": {
			{"//go:build linux", []string{"0:16:Directive"}},
			{"// +build linux", []string{"0:15:Directive"}},
			{"", nil},
			{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
			{"", nil},
			{"import \"embed\"", []string{"0:6:Keyword", "7:14:ImportRef"}},
			{"", nil},
			// A space after the slashes makes it a comment for the go command too.
			{"// go:embed is described here.", []string{"0:30:Comment"}},
			{"//go:embed static", []string{"0:17:Directive"}},
			{"var static embed.FS", []string{"0:3:Keyword", "4:10:Identifier", "11:16:ImportRef", "17:19:TypeRef"}},
			{"", nil},
			{"//export Callback", []string{"0:17:Directive"}},
			{"//line generated.go:10", []string{"0:22:Directive"}},
			{"//golint is not a directive", []string{"0:27:Comment"}},
		},
		"lib.h": {
			{"#ifndef LIB_H", []string{"0:7:Directive", "8:13:Constant"}},
			{"#if defined(_WIN32)", []string{"0:3:Directive", "4:11:Directive", "12:18:Constant"}},
			{"#pragma once", []string{"0:7:Directive"}},
			{"#endif", []string{"0:6:Directive"}},
			{"#endif", []string{"0:6:Directive"}},
		},
		"Program.cs": {
			{"#!/usr/bin/env dotnet-script", []string{"0:28:Directive"}},
			{"#nullable enable", []string{"0:9:Directive"}},
			{"#region Setup", []string{"0:7:Directive"}},
			{"#endregion", []string{"0:10:Directive"}},
		},
		"build.kts": {
			{"#!/usr/bin/env kotlin", []string{"0:21:Directive"}},
			{"println(1) // done", []string{"0:7:Call", "8:9:NumberLiteral", "11:18:Comment"}},
		},
		"run.sh": {
			{"#!/bin/sh", []string{"0:9:Directive"}},
			// Only the first line of a script is a shebang.
			{"#!/bin/bash", []string{"0:11:Comment"}},
		},
		"ci.yaml": {
			{"%YAML 1.2", []string{"0:9:Directive"}},
			{"---", []string{"0:3:Keyword"}},
			{"a: 1", []string{"0:1:Field", "3:4:NumberLiteral"}},
		},
	} {
		t.Run(path, func(t *testing.T) {
			buf, ext := openDoc(t, path, doc.text())
			doc.check(t, highlighterOf(t, buf, ext))
		})
	}
}

// TestGrammarsPrepare checks that every tree-sitter grammar loads and that its
// highlight query compiles. A broken query leaves the language uncolored
// without any other sign of trouble.
func TestGrammarsPrepare(t *testing.T) {
	for _, lang := range languages {
		hl, ok := lang.newHighlighter().(*tsHighlighter)
		if !ok {
			continue
		}
		if err := hl.grammar.prepare(); err != nil {
			t.Errorf("%s: %s", lang.name(), err)
		}
	}
}

// TestLanguageForPath checks that only the recognized file types get a
// highlighter, and that an unknown one is left alone rather than mishandled.
func TestLanguageForPath(t *testing.T) {
	for path, want := range map[string]string{
		"main.go":         "go",
		"a/b/main.GO":     "go",
		"flake.nix":       "nix",
		"schema.CUE":      "cue",
		"package.json":    "json",
		"settings.jsonc":  "jsonc",
		"a/flake.lock":    "json",
		"Cargo.lock":      "",
		"ci.yaml":         "yaml",
		"ci.yml":          "yaml",
		"run.sh":          "shell",
		"a/b/run.BASH":    "shell",
		"init.zsh":        "shell",
		"/home/u/.zshrc":  "shell",
		".bashrc":         "shell",
		".envrc":          "shell",
		"bashrc":          "",
		"main.c":          "c",
		"lib.h":           "cpp",
		"lib.cpp":         "cpp",
		"a/b/Lib.HPP":     "cpp",
		"Program.cs":      "csharp",
		"Main.java":       "java",
		"Main.kt":         "kotlin",
		"init.gradle.kts": "kotlin",
		"index.ts":        "typescript",
		"App.tsx":         "typescriptreact",
		"app.js":          "javascript",
		"App.jsx":         "javascript",
		"src/main.rs":     "rust",
		"init.lisp":       "commonlisp",
		"queries/go.scm":  "query",
		"lib/srfi.scm":    "",
		"notes.md":        "markdown",
		"notes.markdown":  "markdown",
		"":                "",
		"Makefile":        "",
		"query.sql":       "",
		"go":              "",
		"archive.go.bak":  "",
		"dir.go/file.txt": "",
	} {
		got := ""
		if lang := languageForPath(path); lang != nil {
			got = lang.name()
		}
		if got != want {
			t.Errorf("languageForPath(%q) = %q, want %q", path, got, want)
		}
	}
}

func TestNoHighlightForUnknownType(t *testing.T) {
	buf, ext := openDoc(t, "notes.txt", "some text\n")
	if data := buf.ExtensionData(ext.ID()); data != nil {
		t.Errorf("got extension data %#v for an unknown file type, want none", data)
	}
	// AfterEdit has to tolerate a buffer it never made data for.
	ext.AfterEdit(nil, buf, nil)
	// So does a buffer that was never passed through MakeBufferData at all.
	ext.AfterEdit(nil, new(editor.Buffer), nil)
}

// TestGoRealSources runs the highlighter over this package's own sources. It is
// a sanity check rather than an exact comparison: every line with real content
// should come back with some highlight, and no span may fall outside its line.
func TestGoRealSources(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			src, err := os.ReadFile(entry.Name())
			if err != nil {
				t.Fatal(err)
			}
			buf, ext := openDoc(t, entry.Name(), string(src))
			hl := highlighterOf(t, buf, ext)

			check := func(stage string) {
				t.Helper()
				for i, line := range buf.Content.Lines() {
					text := line.String()
					spans := hl.SyntaxSpans(i, text)
					for _, s := range spans {
						if s.LineNumber != i || s.Start < 0 || s.End > len(text) || s.Start >= s.End {
							t.Fatalf("%s > line %d %q: span %s is out of range", stage, i, text, s)
						}
					}
					// Any line with a word on it is a keyword, an identifier, a
					// comment or a string, so it has to come back colorized.
					// Lines of pure punctuation such as "}}," legitimately do not.
					if len(spans) == 0 && strings.ContainsFunc(text, unicode.IsLetter) {
						t.Errorf("%s > line %d %q: no highlight", stage, i, text)
					}
				}
			}

			check("initial")

			before := buf.Content.Len()
			insertLine(t, buf, ext, before/2)
			if buf.Content.Len() != before+1 {
				t.Fatalf("content length %d after insert, want %d", buf.Content.Len(), before+1)
			}
			check("after insert")
		})
	}
}

// TestNoHighlightForDirectoryListing covers a directory whose own name ends in
// a recognized extension: the buffer lists file names, not code.
func TestNoHighlightForDirectoryListing(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "notes.md")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}

	ext := new(Integration)
	ext.LogEmbed = logger.Embed(t.Logf)
	var edit editor.Editor
	edit.Extend(ext)
	edit.OpenDir(dir, noopOpener{})

	buf := edit.Top()
	if buf.Path != "notes.md" {
		t.Fatalf("buffer path is %q, want the directory name", buf.Path)
	}
	if data := buf.ExtensionData(ext.ID()); data != nil {
		t.Errorf("got extension data %#v for a directory listing, want none", data)
	}
}

// noopOpener satisfies content.OpenFile for a test that never opens anything.
type noopOpener struct{}

func (noopOpener) OpenFile(string) {}
