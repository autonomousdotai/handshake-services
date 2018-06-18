const configs = require('./configs');
const express = require('express');
const logger = require('morgan');
const bodyParser = require('body-parser');
const Twit = require('twit');
const utils = require('./utils');
var path = require('path');
var fs = require('fs');

const T = new Twit({
	consumer_key:         configs.consumer_key,
	consumer_secret:      configs.consumer_secret,
	access_token:         configs.access_token,
	access_token_secret:  configs.access_token_secret,
	app_only_auth:        true,
	timeout_ms:           60*1000,  // optional HTTP request timeout to apply to all requests.
	strictSSL:            true,     // optional - requires SSL certificates to be valid.
})

app = express();
app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

const router = express.Router();
router.get('/', function(req, res) {
	res.json(utils.responseSuccess('', 'Twitter REST Service'));
});

router.get('/:username?', function(req, res) {
	if (req.params.username) {
		try {
			let arr = readFile('follows.txt');
			console.log(arr);
			let index = arr.indexOf(req.params.username.toLowerCase());
		  	if (index > -1) {
				res.json(utils.responseSuccess(true));
			}
			else{
				T.get('followers/list', { screen_name: 'ninja_org', count: '10', cursor: -1 },  function (err, data, response) {
					let result = false;
					if (data && data.users && data.users.length > 0) {
						console.log(data.users.length);
						for (let u of data.users) {
							if(u.screen_name == req.params.username){
								result = true;
								arr.push(u.screen_name);
								saveFile("follows.txt", arr);
								console.log('found', u);
								break;
							}
						}
					}
					res.json(utils.responseSuccess(result));
				})
			}	
		} catch (e) {
			console.log('Error: ', e.message);
			res.json(utils.responseError('Error occured, please try again'));
		}
	}
	else{
		res.json(utils.responseError('Invalid params'));
	}
});

function saveFile(fileName, data) {
  let filePath =  path.resolve(__dirname, fileName);
		
		fs.writeFileSync(filePath, data, {flag: 'w'}, function(err) {
			if(err) {
				console.log(err);
			}
		});
}

function readFile(fileName) {
	let filePath =  path.resolve(__dirname,  fileName);
	let str = fs.readFileSync(filePath).toString();
	if(str){
		return str.split(",");
	}
	else{
		return [];
	}
}

app.use('/', router);

app.listen(configs.port, function() {
	console.log('Twitter service listening on port ' + configs.port + '!')
});
