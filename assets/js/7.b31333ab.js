(window.webpackJsonp=window.webpackJsonp||[]).push([[7],{362:function(t,s,a){t.exports=a.p+"assets/img/README-lib-architecture.drawio.0c9c97e5.svg"},374:function(t,s,a){"use strict";a.r(s);var n=a(44),e=Object(n.a)({},(function(){var t=this,s=t.$createElement,n=t._self._c||s;return n("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[n("h1",{attrs:{align:"center"}},[n("img",{attrs:{src:"https://user-images.githubusercontent.com/88447/67148586-4010a580-f2a1-11e9-9ece-91acf37b0c6f.png",alt:"nimona",width:"250px"}})]),t._v(" "),n("h4",{attrs:{align:"center"}},[t._v("a new internet stack; or something like it.")]),t._v(" "),n("p",{attrs:{align:"center"}},[n("a",{attrs:{href:"https://github.com/nimona/go-nimona/actions"}},[n("img",{attrs:{src:"https://github.com/nimona/go-nimona/workflows/CI/badge.svg?style=flat",alt:"Actions Status"}})]),t._v(" "),n("a",{attrs:{href:"https://codeclimate.com/github/nimona/go-nimona"}},[n("img",{attrs:{src:"https://img.shields.io/codeclimate/coverage/nimona/go-nimona",alt:"Coverage"}})]),t._v(" "),n("a",{attrs:{href:"https://github.com/nimona/go-nimona/commits/main"}},[n("img",{attrs:{src:"https://img.shields.io/github/last-commit/nimona/go-nimona.svg?style=flat&logo=github&logoColor=white",alt:"GitHub last commit"}})]),t._v(" "),n("a",{attrs:{href:"https://github.com/nimona/go-nimona/issues"}},[n("img",{attrs:{src:"https://img.shields.io/github/issues-raw/nimona/go-nimona.svg?style=flat&logo=github&logoColor=white",alt:"GitHub issues"}})]),t._v(" "),n("a",{attrs:{href:"https://github.com/nimona/go-nimona/pulls"}},[n("img",{attrs:{src:"https://img.shields.io/github/issues-pr-raw/nimona/go-nimona.svg?style=flat&logo=github&logoColor=white",alt:"GitHub pull requests"}})]),t._v(" "),n("a",{attrs:{href:"https://github.com/nimona/go-nimona/blob/main/LICENSE"}},[n("img",{attrs:{src:"https://img.shields.io/github/license/nimona/go-nimona.svg?style=flat",alt:"License Status"}})])]),t._v(" "),n("hr"),t._v(" "),n("p",[n("strong",[t._v("WARNING")]),t._v(": Nimona is still in its very early stages of design and development\nand will stay like this for a while."),n("br"),t._v("\nDocumentation is slowly starting to pop up, but everything is still pretty much\nin flux.")]),t._v(" "),n("hr"),t._v(" "),n("h1",{attrs:{id:"nimona"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#nimona"}},[t._v("#")]),t._v(" Nimona")]),t._v(" "),n("p",[t._v("Nimona’s main goal is to provide a number of layers/components to help with\nthe challenges presented when dealing with decentralized and peer to peer\napplications.")]),t._v(" "),n("h2",{attrs:{id:"development"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#development"}},[t._v("#")]),t._v(" Development")]),t._v(" "),n("h3",{attrs:{id:"requirements"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#requirements"}},[t._v("#")]),t._v(" Requirements")]),t._v(" "),n("ul",[n("li",[t._v("Go 1.15+ with modules enabled")]),t._v(" "),n("li",[t._v("Make")])]),t._v(" "),n("h3",{attrs:{id:"getting-started"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#getting-started"}},[t._v("#")]),t._v(" Getting Started")]),t._v(" "),n("div",{staticClass:"language- extra-class"},[n("pre",{pre:!0,attrs:{class:"language-text"}},[n("code",[t._v("git clone https://github.com/nimona/go-nimona.git go-nimona\ncd go-nimona\nmake deps\n")])])]),n("h3",{attrs:{id:"process-workflow"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#process-workflow"}},[t._v("#")]),t._v(" Process / Workflow")]),t._v(" "),n("p",[t._v("Nimona is developed using "),n("a",{attrs:{href:"https://commonflow.org/",target:"_blank",rel:"noopener noreferrer"}},[t._v("Git Common-Flow"),n("OutboundLink")],1),t._v(", which is\nessentially "),n("a",{attrs:{href:"http://scottchacon.com/2011/08/31/github-flow.html",target:"_blank",rel:"noopener noreferrer"}},[t._v("GitHub Flow"),n("OutboundLink")],1),t._v("\nwith the addition of versioned releases, and optional release branches.")]),t._v(" "),n("p",[t._v("In addition to the Common-Flow spec, contributors are also highly encouraged to\n"),n("a",{attrs:{href:"https://git-scm.com/book/en/v2/Git-Tools-Signing-Your-Work",target:"_blank",rel:"noopener noreferrer"}},[t._v("sign commits"),n("OutboundLink")],1),t._v(".")]),t._v(" "),n("h2",{attrs:{id:"library-architecture"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#library-architecture"}},[t._v("#")]),t._v(" Library Architecture")]),t._v(" "),n("p",[n("img",{attrs:{src:a(362),alt:"Library Architecture"}})]),t._v(" "),n("h3",{attrs:{id:"network"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#network"}},[t._v("#")]),t._v(" Network")]),t._v(" "),n("p",[t._v("Package "),n("code",[t._v("exchange")]),t._v(" is responsible for a number of things around connections and\nobject exchange, as well as relaying objects to inaccessible peers.")]),t._v(" "),n("div",{staticClass:"language-go extra-class"},[n("pre",{pre:!0,attrs:{class:"language-go"}},[n("code",[n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("type")]),t._v(" Network "),n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("interface")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("Subscribe")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\n        filters "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("...")]),t._v("EnvelopeFilter"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" EnvelopeSubscription\n    "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("Send")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\n        ctx context"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Context"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n        object object"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Object"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n        recipient "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("peer"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("ConnectionInfo"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("error")]),t._v("\n    "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("Listen")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\n        ctx context"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Context"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n        bindAddress "),n("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("Listener"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("error")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),n("h3",{attrs:{id:"resolver"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#resolver"}},[t._v("#")]),t._v(" Resolver")]),t._v(" "),n("p",[t._v("Package "),n("code",[t._v("resolver")]),t._v(" is responsible for looking up peers on the network that\nfulfill specific requirements.")]),t._v(" "),n("div",{staticClass:"language-go extra-class"},[n("pre",{pre:!0,attrs:{class:"language-go"}},[n("code",[n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("type")]),t._v(" Resolver "),n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("interface")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("Lookup")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\n        ctx context"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Context"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n        opts "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("...")]),t._v("LookupOption"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<-")]),n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("chan")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("peer"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("ConnectionInfo"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("error")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),n("p",[t._v("The currently available "),n("code",[t._v("LookupOption")]),t._v(" are the following, and can be used\non their own or in groups.")]),t._v(" "),n("div",{staticClass:"language-go extra-class"},[n("pre",{pre:!0,attrs:{class:"language-go"}},[n("code",[n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("LookupByCID")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("hash object"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("CID"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" LookupOption "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("...")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("LookupByContentType")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("contentType "),n("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" LookupOption "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("...")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("LookupByIdentity")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("key crypto"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("PublicKey"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" LookupOption "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("...")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("LookupByCertificateSigner")]),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("key crypto"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("PublicKey"),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" LookupOption "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v("...")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])])])}),[],!1,null,null,null);s.default=e.exports}}]);