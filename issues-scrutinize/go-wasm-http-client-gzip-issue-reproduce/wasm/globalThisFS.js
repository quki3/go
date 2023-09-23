
function globalThisFS() {
	if(!globalThis.fs){
		console.log("globalThis.fs = false | then create globalThis.fs()");
		/*---creating globalThis.fs---*/

	}else{  console.log("globalThis.fs = true")}
}
