
"use strict";
const debuger = (x) =>{
	console.log(`debugger => ${x}`)
}
if(!globalThis.fs){ //globalThis.fs = false?
	debuger("globalThis.fs = false creating ... done!");
		
	let outputBuf = "";
	globalThis.fs = {
		constants:{ 
			O_WRONLY: -1, //i/o constant normaly used to WRITE ONLY 
			O_RDWR: -1,
			O_CREAT: -1,
			O_TRUNC: -1,
			O_APPEND: -1,
			O_EXCL: -1 
		}

	};
}


