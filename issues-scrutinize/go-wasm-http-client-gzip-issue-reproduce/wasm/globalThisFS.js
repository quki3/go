
(()=>{
	if(!globalThis.fs){ //globalThis.fs = false?
		console.log(`globalThis.fs = false |
		then creating globalThis.fs()`);
		/*---creating globalThis.fs---*/
		let outputBuf = "";
		globalThis.fs = {
			constants: O_WRONLY:-1
		}
		console.log(globalThis.fs)

	}else{  console.log("globalThis.fs = true")}
})()
