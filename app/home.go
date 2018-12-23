package app

import (
	"github.com/rajnikant12345/engine"
	"syscall/js"
)

func CreateHome()*engine.Element {
	e := engine.NewElement("div")
	e.SetInnerHtml(`
<article class="athelas">
  <div class="center measure-wide f3 lh-copy ph2">
    <h1 class="f3 lh-title">ABOUT</h1>
    <p class="f5 dark-gray fw2 tracked">
		I am a senior developer with 10 years of software development experience. I like working on versatile set of problems.
		I have experience in handling and delivering complex project with end to end functionality. My experience consisits of
		working in different organizationas and tackling different set of problems.
    </p>
	<h1 class="f3 lh-title">Open Source Contributions</h1>
    <div class="f5 dark-gray fw2 tracked">
		<ul >
			<li>
  			<a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" href="https://github.com/gfto/toybox/blob/master/toys/pending/dhcp6.c">Contributed DHCP6 client for toybox</a>
			</li>
			<li>
  			<a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" href="https://github.com/gfto/toybox/blob/master/toys/pending/dhcp6.c">Contributed IP-Link and IP-Routing and IP-Neighbour for toybox</a>
			</li>
    </ul> 
	</div>
	<h1 class="f3 lh-title">My Strengths</h1>
	<div class="f5 dark-gray fw2 tracked">
	<ul >
  		<li>Excellent problem solving skills.</li>
  		<li>Ability to adapt.</li>
  		<li>Ability to debug.</li>
		<li>Ability to design.</li>
  		<li>I can work as, individual contributor and as a team contributor.</li>
  		<li>I can autonomously drive projects and deliver an excellent solutions.</li>
		<li>I can learn any programming language and code.</li>
  		<li>I love designing solutions.</li>
  		<li>One of my strength is designing, flexible, scalable and high performance system.</li>
	</ul> 
	</div>
  </div>
</article>`)
	return e
}


func CreateContact()*engine.Element {
	e := engine.NewElement("div")
	e.SetInnerHtml("Contact")
	return e
}



func CreateSkills()*engine.Element {
	e := engine.NewElement("div")
	e.SetInnerHtml(`
<article class="athelas">
  <div class="center measure-wide f3 lh-copy ph2">
    <h1 class="f3 lh-title">Key Skills</h1>
	<div class="f5 dark-gray fw2 tracked">
	<ul >
  		<li>C++/C, Golang ( GO )</li>
  		<li>Microservices, Docker, Docker Compose , Distributed Systems</li>
  		<li>Information Security and Data Protection</li>
		<li>Crypto Algorithms (RSA, AES, ECC), PKI, Secure Applications</li>
  		<li>Microsoft Crypto Service Providers , Microsoft Crypto Next Generation Providers, OpenSSL Engines</li>
  		<li>Operating Systems: Linux, Windows</li>
		<li>Multithreading, Concurrent and Parallel Programming</li>
  		<li>Network Programming – TCP, UDP, Netlink Sockets</li>
  		<li>Design Patterns, OOPS</li>
		<li>Data structures, Algorithms</li>
  		<li>PostgreSQL</li>
	</ul> 
	</div>
  </div>
</article>`)
	return e
}

func Kylo(i []js.Value) {

	Body.RemoveChild()
	e := engine.NewElement("div")
	e.SetInnerHtml(`
<article class="athelas">
  <div class="center measure-wide f3 lh-copy ph2">
    <h1 class="f3 lh-title">Next Generation Key Management Solution</h1>
    <p class="f5 dark-gray fw2 tracked">
		<a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" href="https://safenet.gemalto.com/data-encryption/enterprise-key-management/virtual-key-secure">This a server solution which can be integrated any application or cloud environment for facilitating key management for encryption. This is a microservice based architecture and
        it is developed mostly using Golang and C. I am working on multiple microservices, and completely responsible for designing and developeing Network Attached Encryption.</a>
    </p>
	<h1 class="f3 lh-title">Technlogies Used</h1>
	<div class="f5 dark-gray fw2 tracked">
	<ul >
  		<li>C, GO</li>
  		<li>Microservices</li>
  		<li>Docker, Docker-Compose</li>
		<li>Network Programming</li>
		<li>REST API</li>
		<li>Network Attached Encryption</li>
		<li>The Key Management Interoperability Protocol (KMIP) </li>
		<li>AES, RSA, ECC</li>
		<li>PostgreSQL</li>
	</ul> 
	</div>
  </div>
</article>`)
	Body.AddChild(e)
}

func PdbCtl(i []js.Value) {
	Body.RemoveChild()
	e := engine.NewElement("div")
	e.SetInnerHtml(`
<article class="athelas">
  <div class="center measure-wide f3 lh-copy ph2">
    <h1 class="f3 lh-title">Database Migration Utility</h1>
    <p class="f5 dark-gray fw2 tracked">
		This is a utilty which wraps up the an existing product ProtectDB by Gemalto. This utility is a client which integrates the ProtectDB database migration path with
		Next Generation Key Secure.
    </p>
	<h1 class="f3 lh-title">Technlogies Used</h1>
	<div class="f5 dark-gray fw2 tracked">
	<ul >
  		<li>C, GO</li>
  		<li>Docker, Docker-Compose</li>
		<li>Network Programming</li>
	</ul> 
	</div>
  </div>
</article>`)
	Body.AddChild(e)
}

func RegisterCallbacks() {
	println("callback registered")
	js.Global().Set("kylo", js.NewCallback(Kylo))
	js.Global().Set("pdbctl", js.NewCallback(PdbCtl))
}



func CreateOrgs()*engine.Element {

	e := engine.NewElement("div")
	e.SetInnerHtml(`<article class="athelas">
	<div class="center measure-wide f3 lh-copy ph2">
	<h1 class="f3 lh-title">Gemlato Noida ( April 2016 ---- Present )</h1>
	<div class="f5 dark-gray fw2 tracked nested-list-reset">
	<ul >
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="kylo();">Next Generation Key Management Solution</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="pdbctl();">Database Migration Utility</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="cng();">Microsoft CNG(Cryptography API: Next Generation)</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="csp();">Microsoft CSP(Cryptographic Service Provider)</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="openssl();">OPENSSL engine for EToken</a></li>
	</ul>
	</div>
	</div>

	<div class="center measure-wide f3 lh-copy ph2">
	<h1 class="f3 lh-title">Samsung Noida ( Feb 2013 ---- March 2016 )</h1>
	<div class="f5 dark-gray fw2 tracked nested-list-reset">
	<ul >
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray">SIEM Tool for Samsung R&D</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray">ToyBox Project</a></li>
	</ul>
	</div>
	</div>

	<div class="center measure-wide f3 lh-copy ph2">
	<h1 class="f3 lh-title">Symantec Pune ( Jan 2011 ---- Feb 2013 )</h1>
	<div class="f5 dark-gray fw2 tracked nested-list-reset">
	<ul >
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray">Content Scanning API ( Enterprise Security )</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray">Onboad Antivirus Scanning NetApp</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray">Data Collection SOAP client</a></li>
	</ul>
	</div>
	</div>

	


	</article>`)
	return e
}
