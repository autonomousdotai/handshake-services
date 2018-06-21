var fs = require('fs');
let localConfig = {};
let envConfig = {};
if (fs.existsSync(__dirname + '/config-local.js')) {
  localConfig = require('./config-local');
}

const baseConfig = {
    consumer_key:         '<consumer key>',
    consumer_secret:      '<consumer secret>',
    access_token:         '<access token>',
    access_token_secret:  '<access token secret>',
    port: 5000,
}

const envs = {
    'staging': 'config-staging',
    'production': 'config-production',
};
const env = process.env.NODE_ENV || 'default';

if (envs[env]) {
    if (fs.existsSync(__dirname + '/' + envs[env] + '.js')) {
        envConfig = require('./' + envs[env]);
    }
}

module.exports = Object.assign({}, baseConfig, envConfig, localConfig);
