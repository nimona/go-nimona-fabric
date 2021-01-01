(window.webpackJsonp=window.webpackJsonp||[]).push([[21],{384:function(t,e,s){"use strict";s.r(e);var a=s(45),n=Object(a.a)({},(function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[s("h1",{attrs:{id:"certificates"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#certificates"}},[t._v("#")]),t._v(" Certificates")]),t._v(" "),s("p",[s("em",[t._v("Note: Work in progress.")])]),t._v(" "),s("p",[t._v("Certificates enable users to authorize applications to act on their behalf\nin the network while limiting the data they can interact with and the actions\nthey can perform.")]),t._v(" "),s("p",[t._v("A certificate is a signed object of type "),s("code",[t._v("nimona.io/object.Certificate")]),t._v(" and\nconsists of the following required attributes:")]),t._v(" "),s("ul",[s("li",[s("code",[t._v("subject")]),t._v(" (string): The public part of the application's key pair.")]),t._v(" "),s("li",[s("code",[t._v("resources")]),t._v(" (repeated string): Resources the application has been permitted\nto access. Can either be an exact match for an object type, or a glob.")]),t._v(" "),s("li",[s("code",[t._v("actions")]),t._v(" (repeated string): The type of action permitted for each resource.\nThis is a one-to-one matching with the resources mentioned above. The action\nat index 0 is related to the resources at index 0. Actions currently defined\nare only "),s("code",[t._v("read")]),t._v(" and "),s("code",[t._v("create")]),t._v(".")]),t._v(" "),s("li",[s("code",[t._v("created")]),t._v(" (string, ISO-8601): The timestamp the certificate was created.")]),t._v(" "),s("li",[s("code",[t._v("expires")]),t._v(" (string, ISO-8601): The timestamp the certificate expires.")])]),t._v(" "),s("p",[t._v("An application must include a certificate in every object it sends on behalf of\na user.")]),t._v(" "),s("p",[t._v("If a peer receives an object with a certificate that has either expired, or\ndoes not grant the requester the permissions it needs to perform an action,\nthe remote party simply ignores the object or request.")]),t._v(" "),s("div",{staticClass:"language-json extra-class"},[s("pre",{pre:!0,attrs:{class:"language-json"}},[s("code",[s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"type:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"nimona.io/object.Certificate"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"data:m"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"subject:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"ed25519.2h8Qu2TJCpnwv7jUaQLpazsxMW4iCaTAFgxoi5crsEAs"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"resources:as"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"nimona.io/profile.Profile"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"nimona.io/profile.ProfileRequest"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"mochi.io/conversation.*"')]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"actions:as"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"create"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"create"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"read"')]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"created:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"2020-06-25T19:16:43Z"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"expires:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"2021-06-25T19:16:43Z"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"_signature:m"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"signer:o"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"ed25519.8mE4CeLLCwpyfqyNFkT6gV32ZYcYP6jt1yzMDmzbxxRL"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"alg:o"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"OH_ES256"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"x:d"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"x0..."')]),t._v("\n  "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),s("h2",{attrs:{id:"certificate-requests"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#certificate-requests"}},[t._v("#")]),t._v(" Certificate Requests")]),t._v(" "),s("p",[t._v("Applications need to request a certification from the user.\nTo do that they create, sign, and give to the user a certificate request.")]),t._v(" "),s("p",[t._v("The user needs to load the certificate request using an identity application\nthat manages their identity keys, verify that they are happy with the\npermissions the application is asking for, create a certificate and sign it.")]),t._v(" "),s("p",[t._v("The certificate request can optionally include a request for a profile from\nthe user.")]),t._v(" "),s("div",{staticClass:"language-json extra-class"},[s("pre",{pre:!0,attrs:{class:"language-json"}},[s("code",[s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"type:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"nimona.io/object.CertificateRequest"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"data:m"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"applicationName:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"Foobar"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"applicationDescription:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"An app that does nothing"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"applicationURL:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"https://github.com/nimona"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"applicationIcon:d"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"x0..."')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"applicationBanner:d"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"x0..."')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"requestProfile:b"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("true")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"certificateSubject:s"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"ed25519.2h8Qu2TJCpnwv7jUaQLpazsxMW4iCaTAFgxoi5crsEAs"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"certificateResources:as"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"nimona.io/profile.Profile"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"nimona.io/profile.ProfileRequest"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"mochi.io/conversation.*"')]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token property"}},[t._v('"certificateActions:as"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"create"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"create"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n      "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"read"')]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n  "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),s("p",[t._v("The certificate request can be currently provided in two ways:")]),t._v(" "),s("p",[t._v("a. Through the use of a QR code that the user must scan using their identity\napplication.\nThis is mainly used for when the application is running on a device where the\nuser does not have an identity application installed.")]),t._v(" "),s("p",[t._v("Once the certificate has been created and signed, the identity application\nwill lookup the certificate request's signer in the network and directly\nsend them the certificate.")]),t._v(" "),s("p",[t._v("b. Through a link with a custom URL using the "),s("code",[t._v("nimona://")]),t._v(" scheme and the\n"),s("code",[t._v("certificate-request")]),t._v(" host.\nThe certificate request itself must be provided via a query parameter named\n"),s("code",[t._v("certificateRequest")]),t._v(".")]),t._v(" "),s("p",[t._v("An optional query param "),s("code",[t._v("returnPath")]),t._v(" can be set in order to define how the\nidentity application is expected to return the certificate back to the\nrequester.")]),t._v(" "),s("p",[t._v("If not set, the identity app will lookup the requester on the network and send\nthem the certificate.")]),t._v(" "),s("p",[t._v("If the "),s("code",[t._v("returnPath")]),t._v(" is set to a URL with either HTTPS or custom scheme, the\nidentity application will append "),s("code",[t._v("&certificate=xxx")]),t._v(" to the return path and\nredirect the user to it.")]),t._v(" "),s("h2",{attrs:{id:"revocation"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#revocation"}},[t._v("#")]),t._v(" Revocation")]),t._v(" "),s("p",[t._v("Currently certificate revocation is not supported, we are working hard on that.")])])}),[],!1,null,null,null);e.default=n.exports}}]);