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


	<h1 class="f3 lh-title">Education</h1>
    <div class="f5 dark-gray fw2 tracked">
	<ul >
  		<li>B.E. Electronics and Communication (Kumaon Enginerring College 2006 Passout) 67%</li>
  		<li>XII CBSE, DAV Merrut, 73.8%</li>
  		<li>X CBSE, DAV Meerut, 67.6%</li>
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

func Cng(i []js.Value) {
	Body.RemoveChild()
	e := engine.NewElement("div")
	e.SetInnerHtml(`
<article class="athelas">
  <div class="center measure-wide f3 lh-copy ph2">
    <h1 class="f3 lh-title">CSP and CNG With Key Secure</h1>
    <p class="f5 dark-gray fw2 tracked">
		Used Microsoft SDK for CSP and CNG and developed an interface with Safenet Keysecure. Now this CSP and CNG can be used with tools like, SignTool, ADCS, ADRMS, IIS all these
		software use key for encryption or signing and Using Safenet CNG and CSP, key always resides in KeySecure.
    </p>
	<h1 class="f3 lh-title">Technlogies Used</h1>
	<div class="f5 dark-gray fw2 tracked">
	<ul >
  		<li>C, C++</li>
  		<li>Windows</li>
		<li>CSP, CNG</li>
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
	js.Global().Set("cng", js.NewCallback(Cng))
}



func CreateOrgs()*engine.Element {

	e := engine.NewElement("div")
	e.SetInnerHtml(`<article class="athelas">
	<div class="center measure-wide f3 lh-copy ph2">
	<h1 class="f3 lh-title">Gemlato Noida ( April 2016 ---- Present )</h1>
	<div class="f5 dark-gray fw2 tracked">
	<ul >
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="kylo();">Next Generation Key Management Solution</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="pdbctl();">Database Migration Utility</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="cng();">Microsoft CNG(Cryptography API: Next Generation)</a></li>
	<li><a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" onClick="cng();">Microsoft CSP(Cryptographic Service Provider)</a></li>
	<li>OPENSSL engine for EToken (C , C++, OpenSSL, Linux, PKCS11)</a></li>
	</ul>
	</div>
	</div>

	<div class="center measure-wide f3 lh-copy ph2">
	<h1 class="f3 lh-title">Samsung Noida ( Feb 2013 ---- March 2016 )</h1>
	<div class="f5 dark-gray fw2 tracked">
	<ul>
	<li>Development of Open source command s for Samsung, ip link, ip route, ip neighbour, ftpd, dhcp server. (C, Sockets, Open Source, Linux)</li>
	<li>Development of data collection framework, for Samsung products. Collects data from different Samsung products and send it to a common server. (C++, Linux, Windows, Tizen, Design Patterns)</li>
	</ul>
	</div>
	</div>

	<div class="center measure-wide f3 lh-copy ph2">
	<h1 class="f3 lh-title">Symantec, Pune ( Jan 2011 ---- Feb 2013 )</h1>
	<div class="f5 dark-gray fw2 tracked ">
	<ul >
	<li>Written SOAP client for Symantec Framework for Application data collection. (C++, SOAP, Linux, Windows)</li>
	<li>Worked with NetApp on on-board Antivirus product. (Linux, C++)</li>
	<li>Integrated Symantec micro definition system in On-board AV and scanning component.  (Linux, Networking, Curl, Windows, C++)</li>
	<li>Individual contributor in Content Scanning API, part of Symantec Scan Engine.  (C++ Linux, Windows, Solaris)</li>
	<li>Support different Symantec products for integration with content Scanning API, in both multithreaded and multiprocess scenario.</li>
	</ul>
	</div>
	</div>

	<div class="center measure-wide f3 lh-copy ph2">
	<h1 class="f3 lh-title">Persistent Systems, Pune ( April 2010 ---- Dec 2010)</h1>
	<div class="f5 dark-gray fw2 tracked ">
	<ul >
	<li>Data manipulation component support. This was a support kind of profile in which, I use to fix issues on the product of Agilent. (C++, Linux)</li>
	</ul>
	</div>
	</div>

	<div class="center measure-wide f3 lh-copy ph2">
	<h1 class="f3 lh-title">Clear-Trail , Noida ( August 2008 ---- Feb 2010)</h1>
	<div class="f5 dark-gray fw2 tracked">
	<ul >
	<li>Developed Passive Lawful Interception system (Call and SMS). Used for running scan on existing data from STM1 link and trace out required data. (C++, Linux, Windows, Socket Programming, Design Patterns)</li>
	<li>Active Lawful Interception System (Call and SMS). Used for gathering real-time data based on query and record the call based on that. (C++, Linux, Windows, Socket Programming, Design Patterns)</li>
	</ul>
	</div>
	</div>



	</article>`)
	return e
}
