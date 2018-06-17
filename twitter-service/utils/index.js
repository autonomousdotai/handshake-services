const responseError = function(error) {
	return {
		status: 0,
		message: error,
	}
}

const responseSuccess = function(data = '', message = '') {
	const resp = {
		status: 1,
		data: data,
	};
	if (data) {
		resp.data = data;
	}
	if (message) {
		resp.message = message;
	}
	return resp;
}

module.exports = {
    responseError,
    responseSuccess,
}