package icons

// IconExt is used to represent an icon with a specific extension.
var IconExt = map[string]*IconInfo{
	"htm":               IconSet["html"],
	"html":              IconSet["html"],
	"xhtml":             IconSet["html"],
	"html_vm":           IconSet["html"],
	"asp":               IconSet["html"],
	"jade":              IconSet["pug"],
	"pug":               IconSet["pug"],
	"md":                IconSet["markdown"],
	"markdown":          IconSet["markdown"],
	"rst":               IconSet["markdown"],
	"blink":             IconSet["blink"],
	"css":               IconSet["css"],
	"scss":              IconSet["sass"],
	"sass":              IconSet["sass"],
	"less":              IconSet["less"],
	"json":              IconSet["json"],
	"tsbuildinfo":       IconSet["json"],
	"json5":             IconSet["json"],
	"jsonl":             IconSet["json"],
	"ndjson":            IconSet["json"],
	"jinja":             IconSet["jinja"],
	"jinja2":            IconSet["jinja"],
	"j2":                IconSet["jinja"],
	"jinja-html":        IconSet["jinja"],
	"sublime-project":   IconSet["sublime"],
	"sublime-workspace": IconSet["sublime"],
	"yaml":              IconSet["yaml"],
	"yaml-tmlanguage":   IconSet["yaml"],
	"yml":               IconSet["yaml"],
	"xml":               IconSet["xml"],
	"plist":             IconSet["xml"],
	"xsd":               IconSet["xml"],
	"dtd":               IconSet["xml"],
	"xsl":               IconSet["xml"],
	"xslt":              IconSet["xml"],
	"resx":              IconSet["xml"],
	"iml":               IconSet["xml"],
	"tmLanguage":        IconSet["xml"],
	"manifest":          IconSet["xml"],
	"project":           IconSet["xml"],
	"png":               IconSet["image"],
	"jpeg":              IconSet["image"],
	"jpg":               IconSet["image"],
	"gif":               IconSet["image"],
	"ico":               IconSet["image"],
	"tif":               IconSet["image"],
	"tiff":              IconSet["image"],
	"psd":               IconSet["image"],
	"psb":               IconSet["image"],
	"ami":               IconSet["image"],
	"apx":               IconSet["image"],
	"bmp":               IconSet["image"],
	"bpg":               IconSet["image"],
	"brk":               IconSet["image"],
	"cur":               IconSet["image"],
	"dds":               IconSet["image"],
	"dng":               IconSet["image"],
	"exr":               IconSet["image"],
	"fpx":               IconSet["image"],
	"gbr":               IconSet["image"],
	"img":               IconSet["image"],
	"jbig2":             IconSet["image"],
	"jb2":               IconSet["image"],
	"jng":               IconSet["image"],
	"jxr":               IconSet["image"],
	"pbm":               IconSet["image"],
	"pgf":               IconSet["image"],
	"pic":               IconSet["image"],
	"raw":               IconSet["image"],
	"webp":              IconSet["image"],
	"eps":               IconSet["image"],
	"afphoto":           IconSet["image"],
	"ase":               IconSet["image"],
	"aseprite":          IconSet["image"],
	"clip":              IconSet["image"],
	"cpt":               IconSet["image"],
	"heif":              IconSet["image"],
	"heic":              IconSet["image"],
	"kra":               IconSet["image"],
	"mdp":               IconSet["image"],
	"ora":               IconSet["image"],
	"pdn":               IconSet["image"],
	"reb":               IconSet["image"],
	"sai":               IconSet["image"],
	"tga":               IconSet["image"],
	"xcf":               IconSet["image"],
	"js":                IconSet["javascript"],
	"esx":               IconSet["javascript"],
	"mj":                IconSet["javascript"],
	"jsx":               IconSet["react"],
	"tsx":               IconSet["react_ts"],
	"ini":               IconSet["settings"],
	"dlc":               IconSet["settings"],
	"dll":               IconSet["settings"],
	"config":            IconSet["settings"],
	"conf":              IconSet["settings"],
	"properties":        IconSet["settings"],
	"prop":              IconSet["settings"],
	"settings":          IconSet["settings"],
	"option":            IconSet["settings"],
	"props":             IconSet["settings"],
	"toml":              IconSet["settings"],
	"prefs":             IconSet["settings"],
	"dotsettings":       IconSet["settings"],
	"cfg":               IconSet["settings"],
	"ts":                IconSet["typescript"],
	"marko":             IconSet["markojs"],
	"pdf":               IconSet["pdf"],
	"xlsx":              IconSet["table"],
	"xls":               IconSet["table"],
	"csv":               IconSet["table"],
	"tsv":               IconSet["table"],
	"vscodeignore":      IconSet["vscode"],
	"vsixmanifest":      IconSet["vscode"],
	"vsix":              IconSet["vscode"],
	"code-workplace":    IconSet["vscode"],
	"csproj":            IconSet["visualstudio"],
	"ruleset":           IconSet["visualstudio"],
	"sln":               IconSet["visualstudio"],
	"suo":               IconSet["visualstudio"],
	"vb":                IconSet["visualstudio"],
	"vbs":               IconSet["visualstudio"],
	"vcxitems":          IconSet["visualstudio"],
	"vcxproj":           IconSet["visualstudio"],
	"pdb":               IconSet["database"],
	"sql":               IconSet["mysql"],
	"pks":               IconSet["database"],
	"pkb":               IconSet["database"],
	"accdb":             IconSet["database"],
	"mdb":               IconSet["database"],
	"sqlite":            IconSet["sqlite"],
	"sqlite3":           IconSet["sqlite"],
	"pgsql":             IconSet["postgresql"],
	"postgres":          IconSet["postgresql"],
	"psql":              IconSet["postgresql"],
	"cs":                IconSet["csharp"],
	"csx":               IconSet["csharp"],
	"qs":                IconSet["qsharp"],
	"zip":               IconSet["zip"],
	"tar":               IconSet["zip"],
	"gz":                IconSet["zip"],
	"xz":                IconSet["zip"],
	"br":                IconSet["zip"],
	"bzip2":             IconSet["zip"],
	"gzip":              IconSet["zip"],
	"brotli":            IconSet["zip"],
	"7z":                IconSet["zip"],
	"rar":               IconSet["zip"],
	"tgz":               IconSet["zip"],
	"vala":              IconSet["vala"],
	"zig":               IconSet["zig"],
	"exe":               IconSet["exe"],
	"msi":               IconSet["exe"],
	"java":              IconSet["java"],
	"jar":               IconSet["java"],
	"jsp":               IconSet["java"],
	"c":                 IconSet["c"],
	"m":                 IconSet["c"],
	"i":                 IconSet["c"],
	"mi":                IconSet["c"],
	"h":                 IconSet["h"],
	"cc":                IconSet["cpp"],
	"cpp":               IconSet["cpp"],
	"cxx":               IconSet["cpp"],
	"c++":               IconSet["cpp"],
	"cp":                IconSet["cpp"],
	"mm":                IconSet["cpp"],
	"mii":               IconSet["cpp"],
	"ii":                IconSet["cpp"],
	"hh":                IconSet["hpp"],
	"hpp":               IconSet["hpp"],
	"hxx":               IconSet["hpp"],
	"h++":               IconSet["hpp"],
	"hp":                IconSet["hpp"],
	"tcc":               IconSet["hpp"],
	"inl":               IconSet["hpp"],
	"go":                IconSet["go"],
	"py":                IconSet["python"],
	"pyc":               IconSet["python-misc"],
	"whl":               IconSet["python-misc"],
	"url":               IconSet["url"],
	"sh":                IconSet["console"],
	"ksh":               IconSet["console"],
	"csh":               IconSet["console"],
	"tcsh":              IconSet["console"],
	"zsh":               IconSet["console"],
	"bash":              IconSet["console"],
	"bat":               IconSet["console"],
	"cmd":               IconSet["console"],
	"awk":               IconSet["console"],
	"fish":              IconSet["console"],
	"ps1":               IconSet["powershell"],
	"psm1":              IconSet["powershell"],
	"psd1":              IconSet["powershell"],
	"ps1xml":            IconSet["powershell"],
	"psc1":              IconSet["powershell"],
	"pssc":              IconSet["powershell"],
	"gradle":            IconSet["gradle"],
	"doc":               IconSet["word"],
	"docx":              IconSet["word"],
	"rtf":               IconSet["word"],
	"cer":               IconSet["certificate"],
	"cert":              IconSet["certificate"],
	"crt":               IconSet["certificate"],
	"pub":               IconSet["key"],
	"key":               IconSet["key"],
	"pem":               IconSet["key"],
	"asc":               IconSet["key"],
	"gpg":               IconSet["key"],
	"woff":              IconSet["font"],
	"woff2":             IconSet["font"],
	"ttf":               IconSet["font"],
	"eot":               IconSet["font"],
	"suit":              IconSet["font"],
	"otf":               IconSet["font"],
	"bmap":              IconSet["font"],
	"fnt":               IconSet["font"],
	"odttf":             IconSet["font"],
	"ttc":               IconSet["font"],
	"font":              IconSet["font"],
	"fonts":             IconSet["font"],
	"sui":               IconSet["font"],
	"ntf":               IconSet["font"],
	"mrf":               IconSet["font"],
	"lib":               IconSet["lib"],
	"bib":               IconSet["lib"],
	"rb":                IconSet["ruby"],
	"erb":               IconSet["ruby"],
	"fs":                IconSet["fsharp"],
	"fsx":               IconSet["fsharp"],
	"fsi":               IconSet["fsharp"],
	"fsproj":            IconSet["fsharp"],
	"swift":             IconSet["swift"],
	"ino":               IconSet["arduino"],
	"dockerignore":      IconSet["docker"],
	"dockerfile":        IconSet["docker"],
	"tex":               IconSet["tex"],
	"sty":               IconSet["tex"],
	"dtx":               IconSet["tex"],
	"ltx":               IconSet["tex"],
	"pptx":              IconSet["powerpoint"],
	"ppt":               IconSet["powerpoint"],
	"pptm":              IconSet["powerpoint"],
	"potx":              IconSet["powerpoint"],
	"potm":              IconSet["powerpoint"],
	"ppsx":              IconSet["powerpoint"],
	"ppsm":              IconSet["powerpoint"],
	"pps":               IconSet["powerpoint"],
	"ppam":              IconSet["powerpoint"],
	"ppa":               IconSet["powerpoint"],
	"webm":              IconSet["video"],
	"mkv":               IconSet["video"],
	"flv":               IconSet["video"],
	"vob":               IconSet["video"],
	"ogv":               IconSet["video"],
	"ogg":               IconSet["video"],
	"gifv":              IconSet["video"],
	"avi":               IconSet["video"],
	"mov":               IconSet["video"],
	"qt":                IconSet["video"],
	"wmv":               IconSet["video"],
	"yuv":               IconSet["video"],
	"rm":                IconSet["video"],
	"rmvb":              IconSet["video"],
	"mp4":               IconSet["video"],
	"m4v":               IconSet["video"],
	"mpg":               IconSet["video"],
	"mp2":               IconSet["video"],
	"mpeg":              IconSet["video"],
	"mpe":               IconSet["video"],
	"mpv":               IconSet["video"],
	"m2v":               IconSet["video"],
	"vdi":               IconSet["virtual"],
	"vbox":              IconSet["virtual"],
	"vbox-prev":         IconSet["virtual"],
	"ics":               IconSet["email"],
	"mp3":               IconSet["audio"],
	"flac":              IconSet["audio"],
	"m4a":               IconSet["audio"],
	"wma":               IconSet["audio"],
	"aiff":              IconSet["audio"],
	"coffee":            IconSet["coffee"],
	"cson":              IconSet["coffee"],
	"iced":              IconSet["coffee"],
	"txt":               IconSet["document"],
	"graphql":           IconSet["graphql"],
	"gql":               IconSet["graphql"],
	"rs":                IconSet["rust"],
	"raml":              IconSet["raml"],
	"xaml":              IconSet["xaml"],
	"hs":                IconSet["haskell"],
	"kt":                IconSet["kotlin"],
	"kts":               IconSet["kotlin"],
	"patch":             IconSet["git"],
	"lua":               IconSet["lua"],
	"clj":               IconSet["clojure"],
	"cljs":              IconSet["clojure"],
	"cljc":              IconSet["clojure"],
	"groovy":            IconSet["groovy"],
	"r":                 IconSet["r"],
	"rmd":               IconSet["r"],
	"dart":              IconSet["dart"],
	"as":                IconSet["actionscript"],
	"mxml":              IconSet["mxml"],
	"ahk":               IconSet["autohotkey"],
	"swf":               IconSet["flash"],
	"swc":               IconSet["swc"],
	"cmake":             IconSet["cmake"],
	"asm":               IconSet["assembly"],
	"a51":               IconSet["assembly"],
	"inc":               IconSet["assembly"],
	"nasm":              IconSet["assembly"],
	"s":                 IconSet["assembly"],
	"ms":                IconSet["assembly"],
	"agc":               IconSet["assembly"],
	"ags":               IconSet["assembly"],
	"aea":               IconSet["assembly"],
	"argus":             IconSet["assembly"],
	"mitigus":           IconSet["assembly"],
	"binsource":         IconSet["assembly"],
	"vue":               IconSet["vue"],
	"ml":                IconSet["ocaml"],
	"mli":               IconSet["ocaml"],
	"cmx":               IconSet["ocaml"],
	"lock":              IconSet["lock"],
	"hbs":               IconSet["handlebars"],
	"mustache":          IconSet["handlebars"],
	"pm":                IconSet["perl"],
	"raku":              IconSet["perl"],
	"hx":                IconSet["haxe"],
	"pp":                IconSet["puppet"],
	"ex":                IconSet["elixir"],
	"exs":               IconSet["elixir"],
	"eex":               IconSet["elixir"],
	"leex":              IconSet["elixir"],
	"erl":               IconSet["erlang"],
	"twig":              IconSet["twig"],
	"jl":                IconSet["julia"],
	"elm":               IconSet["elm"],
	"pure":              IconSet["purescript"],
	"purs":              IconSet["purescript"],
	"tpl":               IconSet["smarty"],
	"styl":              IconSet["stylus"],
	"merlin":            IconSet["merlin"],
	"v":                 IconSet["verilog"],
	"vhd":               IconSet["verilog"],
	"sv":                IconSet["verilog"],
	"svh":               IconSet["verilog"],
	"robot":             IconSet["robot"],
	"sol":               IconSet["solidity"],
	"yang":              IconSet["yang"],
	"mjml":              IconSet["mjml"],
	"tf":                IconSet["terraform"],
	"tfvars":            IconSet["terraform"],
	"tfstate":           IconSet["terraform"],
	"applescript":       IconSet["applescript"],
	"ipa":               IconSet["applescript"],
	"cake":              IconSet["cake"],
	"nim":               IconSet["nim"],
	"nimble":            IconSet["nim"],
	"apib":              IconSet["apiblueprint"],
	"apiblueprint":      IconSet["apiblueprint"],
	"pcss":              IconSet["postcss"],
	"sss":               IconSet["postcss"],
	"todo":              IconSet["todo"],
	"nix":               IconSet["nix"],
	"slim":              IconSet["slim"],
	"http":              IconSet["http"],
	"rest":              IconSet["http"],
	"apk":               IconSet["android"],
	"env":               IconSet["tune"],
	"jenkinsfile":       IconSet["jenkins"],
	"jenkins":           IconSet["jenkins"],
	"log":               IconSet["log"],
	"ejs":               IconSet["ejs"],
	"djt":               IconSet["django"],
	"pot":               IconSet["i18n"],
	"po":                IconSet["i18n"],
	"mo":                IconSet["i18n"],
	"d":                 IconSet["d"],
	"mdx":               IconSet["mdx"],
	"gd":                IconSet["godot"],
	"godot":             IconSet["godot-assets"],
	"tres":              IconSet["godot-assets"],
	"tscn":              IconSet["godot-assets"],
	"azcli":             IconSet["azure"],
	"vagrantfile":       IconSet["vagrant"],
	"cshtml":            IconSet["razor"],
	"vbhtml":            IconSet["razor"],
	"ad":                IconSet["asciidoc"],
	"adoc":              IconSet["asciidoc"],
	"asciidoc":          IconSet["asciidoc"],
	"edge":              IconSet["edge"],
	"ss":                IconSet["scheme"],
	"scm":               IconSet["scheme"],
	"stl":               IconSet["3d"],
	"obj":               IconSet["3d"],
	"ac":                IconSet["3d"],
	"blend":             IconSet["3d"],
	"mesh":              IconSet["3d"],
	"mqo":               IconSet["3d"],
	"pmd":               IconSet["3d"],
	"pmx":               IconSet["3d"],
	"skp":               IconSet["3d"],
	"vac":               IconSet["3d"],
	"vdp":               IconSet["3d"],
	"vox":               IconSet["3d"],
	"svg":               IconSet["svg"],
	"vimrc":             IconSet["vim"],
	"gvimrc":            IconSet["vim"],
	"exrc":              IconSet["vim"],
	"moon":              IconSet["moonscript"],
	"iso":               IconSet["disc"],
	"f":                 IconSet["fortran"],
	"f77":               IconSet["fortran"],
	"f90":               IconSet["fortran"],
	"f95":               IconSet["fortran"],
	"f03":               IconSet["fortran"],
	"f08":               IconSet["fortran"],
	"tcl":               IconSet["tcl"],
	"liquid":            IconSet["liquid"],
	"p":                 IconSet["prolog"],
	"pro":               IconSet["prolog"],
	"coco":              IconSet["coconut"],
	"sketch":            IconSet["sketch"],
	"opam":              IconSet["opam"],
	"dhallb":            IconSet["dhall"],
	"pwn":               IconSet["pawn"],
	"amx":               IconSet["pawn"],
	"dhall":             IconSet["dhall"],
	"pas":               IconSet["pascal"],
	"unity":             IconSet["shaderlab"],
	"nupkg":             IconSet["nuget"],
	"command":           IconSet["command"],
	"dsc":               IconSet["denizenscript"],
	"deb":               IconSet["debian"],
	"rpm":               IconSet["redhat"],
	"snap":              IconSet["ubuntu"],
	"ebuild":            IconSet["gentoo"],
	"pkg":               IconSet["applescript"],
	"openbsd":           IconSet["freebsd"],
	// "ls":                IconSet["livescript"],
	// "re":                IconSet["reason"],
	// "rei":               IconSet["reason"],
	// "cmj":               IconSet["bucklescript"],
	// "nb":                IconSet["mathematica"],
	// "wl":                IconSet["wolframlanguage"],
	// "wls":               IconSet["wolframlanguage"],
	// "njk":               IconSet["nunjucks"],
	// "nunjucks":          IconSet["nunjucks"],
	// "au3":               IconSet["autoit"],
	// "haml":              IconSet["haml"],
	// "feature":           IconSet["cucumber"],
	// "riot":              IconSet["riot"],
	// "tag":               IconSet["riot"],
	// "vfl":               IconSet["vfl"],
	// "kl":                IconSet["kl"],
	// "cfml":              IconSet["coldfusion"],
	// "cfc":               IconSet["coldfusion"],
	// "lucee":             IconSet["coldfusion"],
	// "cfm":               IconSet["coldfusion"],
	// "cabal":             IconSet["cabal"],
	// "rql":               IconSet["restql"],
	// "restql":            IconSet["restql"],
	// "kv":                IconSet["kivy"],
	// "graphcool":         IconSet["graphcool"],
	// "sbt":               IconSet["sbt"],
	// "cr":                IconSet["crystal"],
	// "ecr":               IconSet["crystal"],
	// "cu":                IconSet["cuda"],
	// "cuh":               IconSet["cuda"],
	// "def":               IconSet["dotjs"],
	// "dot":               IconSet["dotjs"],
	// "jst":               IconSet["dotjs"],
	// "pde":               IconSet["processing"],
	// "wpy":               IconSet["wepy"],
	// "hcl":               IconSet["hcl"],
	// "san":               IconSet["san"],
	// "red":               IconSet["red"],
	// "fxp":               IconSet["foxpro"],
	// "prg":               IconSet["foxpro"],
	// "wat":               IconSet["webassembly"],
	// "wasm":              IconSet["webassembly"],
	// "ipynb":             IconSet["jupyter"],
	// "bal":               IconSet["ballerina"],
	// "balx":              IconSet["ballerina"],
	// "rkt":               IconSet["racket"],
	// "bzl":               IconSet["bazel"],
	// "bazel":             IconSet["bazel"],
	// "mint":              IconSet["mint"],
	// "vm":                IconSet["velocity"],
	// "fhtml":             IconSet["velocity"],
	// "vtl":               IconSet["velocity"],
	// "prisma":            IconSet["prisma"],
	// "abc":               IconSet["abc"],
	// "lisp":              IconSet["lisp"],
	// "lsp":               IconSet["lisp"],
	// "cl":                IconSet["lisp"],
	// "fast":              IconSet["lisp"],
	// "svelte":            IconSet["svelte"],
	// "prw":               IconSet["advpl_prw"],
	// "prx":               IconSet["advpl_prw"],
	// "ptm":               IconSet["advpl_ptm"],
	// "tlpp":              IconSet["advpl_tlpp"],
	// "ch":                IconSet["advpl_include"],
	// "4th":               IconSet["forth"],
	// "fth":               IconSet["forth"],
	// "frt":               IconSet["forth"],
	// "iuml":              IconSet["uml"],
	// "pu":                IconSet["uml"],
	// "puml":              IconSet["uml"],
	// "plantuml":          IconSet["uml"],
	// "wsd":               IconSet["uml"],
	// "sml":               IconSet["sml"],
	// "mlton":             IconSet["sml"],
	// "mlb":               IconSet["sml"],
	// "sig":               IconSet["sml"],
	// "fun":               IconSet["sml"],
	// "cm":                IconSet["sml"],
	// "lex":               IconSet["sml"],
	// "use":               IconSet["sml"],
	// "grm":               IconSet["sml"],
	// "imba":              IconSet["imba"],
	// "drawio":            IconSet["drawio"],
	// "dio":               IconSet["drawio"],
	// "sas":               IconSet["sas"],
	// "sas7bdat":          IconSet["sas"],
	// "sashdat":           IconSet["sas"],
	// "astore":            IconSet["sas"],
	// "ast":               IconSet["sas"],
	// "sast":              IconSet["sas"],
}
