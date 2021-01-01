(window.webpackJsonp=window.webpackJsonp||[]).push([[7],{356:function(t,e,s){t.exports=s.p+"assets/img/np003-streams.drawio.b6b06b4a.svg"},382:function(t,e,s){"use strict";s.r(e);var a=s(42),r=Object(a.a)({},(function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"streams"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#streams"}},[t._v("#")]),t._v(" Streams")]),t._v(" "),a("h2",{attrs:{id:"simple-summary"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#simple-summary"}},[t._v("#")]),t._v(" Simple Summary")]),t._v(" "),a("p",[t._v("Streams provide a way of creating complex mutable data structures using\ndirected acyclic graphs made from objects.")]),t._v(" "),a("h2",{attrs:{id:"problem-statement"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#problem-statement"}},[t._v("#")]),t._v(" Problem Statement")]),t._v(" "),a("p",[t._v("While objects on their own are useful for creating permanent content-addressable\ndata structures, there are very few applications where data never get updated.\nThis is where streams come in, they allow developers to create complex\napplications by applying event driven and event sourcing patterns using graphs\nof individually immutable objects.")]),t._v(" "),a("h2",{attrs:{id:"proposal"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#proposal"}},[t._v("#")]),t._v(" Proposal")]),t._v(" "),a("p",[t._v("Objects in a stream form a directed acyclic graph (DAG) by allowing each of the\nobjects to reference others it depends on or knows of. This graph can then be\nserialized into a linear series of objects that can be replayed consistently by\neveryone that has the same representation of the graph.")]),t._v(" "),a("p",[t._v("Streams are identified by the hash of their root object. This means that even\nthough each of their objects is content-addressable; the stream as a whole is\nnot, as its root hash (and thus identifier) does not change when more objects\nare added to the graph.")]),t._v(" "),a("p",[t._v("The benefit of this is that there is no need to find a way to reference the\nstream as it changes. The downside is that you do not really know if you have\nactually  received the whole stream and whether peers are not holding back on\nyou.")]),t._v(" "),a("p",[a("img",{attrs:{src:s(356),alt:"stream"}})]),t._v(" "),a("h2",{attrs:{id:"structure"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#structure"}},[t._v("#")]),t._v(" Structure")]),t._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"type:s"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"stream:nimona.io/kv"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"data:m"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("h2",{attrs:{id:"access-control"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#access-control"}},[t._v("#")]),t._v(" Access control")]),t._v(" "),a("p",[a("em",[t._v("Note: Work in progress.")])]),t._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"type:s"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"nimona.io/profile.Created"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"owner:a"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"f00"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"policy:m"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"subjects:as"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"*"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"resources:as"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"*"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"action:s"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"READ"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"allow:b"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("true")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("h2",{attrs:{id:"hypothetical-roots"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#hypothetical-roots"}},[t._v("#")]),t._v(" Hypothetical roots")]),t._v(" "),a("p",[t._v('As mentioned before, streams are identified by the hash of their root object. In\norder for a peer to find the providers of a stream and get its objects, it must\nat the very least know its identifier. This is usually not an issue as most\ntimes a peer will learn about the existence of a stream from somewhere before\ndeciding to request it. There are some cases though where that might not be the\ncase, especially when looking for something that might be considered relatively\n"well known".')]),t._v(" "),a("p",[t._v("An example of this would be the profile stream of an identity. Let's say we are\nlooking at a blog post that a single author. Unless the blog post somehow\ncontains a link to the author's profile stream, there is no other way to easily\nfind the stream's identifier.")]),t._v(" "),a("p",[t._v("This is where hypothetical roots come in.")]),t._v(" "),a("p",[t._v("A hypothetical root is an object that identifies a stream and can be assumed\nexists given the type of stream and the author that would have created it. This\nallows peers to find streams unique to an identity without having to somehow\nlearn of their existence.")]),t._v(" "),a("p",[t._v("Since the hypothetical root does not contain a policy, the stream starts off as\npublicly accessible but writable only by the author. The author can subsequently\ndecide to restrict the rest of the stream by using a more strict policy.")]),t._v(" "),a("hr"),t._v(" "),a("p",[t._v("Let's go back to our original example of profile streams.")]),t._v(" "),a("p",[t._v("Assuming that peer "),a("code",[t._v("a11")]),t._v(" wants the profile stream for the identity "),a("code",[t._v("f00")]),t._v(", all it\nhas to do is construct the hypothetical root, calculate its hash, and find\nproviders for it on the network.")]),t._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"type:s"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"nimona.io/profile.Created"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"author:s"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"f00"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("The hash of this object is "),a("code",[t._v("oh1.9KQhQ4UGaQPEyUDAAPDmVJCoHnGtJY7Aun4coFATXCYK")]),t._v("\nand the peer can now lookup the providers for this object, and sync the\nremaining stream.")]),t._v(" "),a("hr"),t._v(" "),a("p",[t._v("The NDL for defining hypothetical roots is as follows. Additional objects can be\ndefined in the stream as needed, but the hypothetical root object itself cannot\nhave additional properties.")]),t._v(" "),a("div",{staticClass:"language-ndl extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[t._v("stream nimona.io/profile {\n    hypothetical root object Created { }\n    signed object NameUpdated {\n        nameFirst string\n        nameLast string\n        dependsOn repeated relationship\n    }\n}\n")])])]),a("h2",{attrs:{id:"synchronization"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#synchronization"}},[t._v("#")]),t._v(" Synchronization")]),t._v(" "),a("p",[a("em",[t._v("Note: Work in progress.")])]),t._v(" "),a("div",{staticClass:"language-ndl extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[t._v("    signed object nimona.io/stream.StreamRequest {\n        nonce string\n        leaves repeated nimona.io/object.Hash\n    }\n")])])]),a("div",{staticClass:"language-ndl extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[t._v("    signed object nimona.io/stream.StreamResponse {\n        nonce string\n        children repeated nimona.io/object.Hash\n    }\n")])])]),a("div",{staticClass:"language-ndl extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[t._v("    signed object nimona.io/stream.Announcement {\n        nonce string\n        leaves repeated nimona.io/object.Hash\n    }\n")])])]),a("h2",{attrs:{id:"subscriptions"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#subscriptions"}},[t._v("#")]),t._v(" Subscriptions")]),t._v(" "),a("p",[t._v('Peers can "subscribe" to stream updates by creating and sending subscriptions to\nother peers. A subscription can be used to subscribe on updates to one or more\nstreams using the streams\' root hash and must also specify an expiration time\nfor the subscription.')]),t._v(" "),a("p",[t._v("When a peer receives or creates an update for a stream, they will go through the\nsubscriptions they have received, and notify the relevant peers about the new\nupdates. If the subscriber does not have have access to the stream, no\nnotification will be sent.")]),t._v(" "),a("div",{staticClass:"language-ndl extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[t._v("signed object nimona.io/stream.Subscription {\n    rootHashes nimona.io/object.Hash\n    expiry nimona.io/object.DateTime\n}\n")])])]),a("p",[t._v("Subscriptions can also be added as stream events. This allows identities and\npeers that have write access to a stream to denote their interest in receiving\nupdates about that stream. In this case "),a("code",[t._v("rootHashes")]),t._v(" should be empty and the\nexpiry is optional.")]),t._v(" "),a("h2",{attrs:{id:"references"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#references"}},[t._v("#")]),t._v(" References")]),t._v(" "),a("ul",[a("li",[a("a",{attrs:{href:"https://docs.textile.io/threads/#threads",target:"_blank",rel:"noopener noreferrer"}},[t._v("https://docs.textile.io/threads/#threads"),a("OutboundLink")],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"https://www.streamr.com/docs/streams",target:"_blank",rel:"noopener noreferrer"}},[t._v("https://www.streamr.com/docs/streams"),a("OutboundLink")],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"https://holochain.org",target:"_blank",rel:"noopener noreferrer"}},[t._v("https://holochain.org"),a("OutboundLink")],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"https://github.com/textileio/go-textile/issues/694",target:"_blank",rel:"noopener noreferrer"}},[t._v("https://github.com/textileio/go-textile/issues/694"),a("OutboundLink")],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"https://tuhrig.de/messages-vs-events-vs-commands",target:"_blank",rel:"noopener noreferrer"}},[t._v("https://tuhrig.de/messages-vs-events-vs-commands"),a("OutboundLink")],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"https://www.swirlds.com/downloads/SWIRLDS-TR-2016-01.pdf",target:"_blank",rel:"noopener noreferrer"}},[t._v("https://www.swirlds.com/downloads/SWIRLDS-TR-2016-01.pdf"),a("OutboundLink")],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"https://arxiv.org/pdf/1710.04469.pdf",target:"_blank",rel:"noopener noreferrer"}},[t._v("https://arxiv.org/pdf/1710.04469.pdf"),a("OutboundLink")],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"http://archagon.net/blog/2018/03/24/data-laced-with-history",target:"_blank",rel:"noopener noreferrer"}},[t._v("http://archagon.net/blog/2018/03/24/data-laced-with-history"),a("OutboundLink")],1)])])])}),[],!1,null,null,null);e.default=r.exports}}]);