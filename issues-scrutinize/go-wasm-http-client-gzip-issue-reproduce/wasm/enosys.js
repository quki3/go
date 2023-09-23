
function enosys() {
	const err = new Error("not implemented");
	err.code = "ENOSYS = NOT IMPLEMENTED";
	return err;
};
