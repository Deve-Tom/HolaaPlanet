window.onload=function (){
const fs = require('fs');
const yaml = require('js-yaml');
try {
    const fileContents = fs.readFileSync('./conf.yaml', 'utf8');
    const data = yaml.safeLoad();
    console.log(data);
} catch (e) {
    console.log(e);
}
}