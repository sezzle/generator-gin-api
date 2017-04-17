'use strict';
var yeoman = require('yeoman-generator');
var mkdirp = require('mkdirp');

var chalk = require('chalk');
var yosay = require('yosay');

module.exports = yeoman.Base.extend({
  prompting: function () {
    // Have Yeoman greet the user.
    this.log(yosay(
      'Welcome to the terrific ' + chalk.red('generator-gin-api') + ' generator!'
    ));
    console.log('\n' +
      '+------------------------------------------------+\n' +
      '| G o  G I N | p r o j e c t | g e n e r a t o r |\n' +
      '+------------------------------------------------+\n' +
      '\n');

    return this.prompt([{
      type: 'input',
      name: 'myappName',
      message: 'What is the name (or path ie myName/myProject if on github.com) of your application',
      default: this.appname.replace(/\s+/g, '-')
    }, {
      type: 'input',
      name: 'myrepoUrl',
      message: 'What is the URL repository of your application? (or project base name under $GOPATH/src/)',
      default: 'github.com'
    }]).then(function (answers) {
      this.log('app name', answers.myappName);
      this.myappName = answers.myappName;
      this.myrepoUrl = answers.myrepoUrl;
    }.bind(this));
  },

  buildTreeFolderAndCopyFiles: function () {
    console.log('Generating tree folders');
    var configDir = 'config/';
    var ginDir = 'gin/';
    var gormDir = 'gorm/';
    var deployDir = 'deploy/';

    mkdirp(configDir);
    mkdirp(ginDir);
    mkdirp(gormDir);
    mkdirp(deployDir);

    this.copy('gitignore', '.gitignore');
    this.copy('config.go', configDir + 'config.go');
    this.copy('localConfig.yaml', configDir + 'localConfig.yaml');
    this.copy('localConfig_sample.yaml', configDir + 'localConfig_sample.yaml');
    this.copy('gorm.go', gormDir + 'gorm.go');
    this.copy('user.go', gormDir + 'user.go');
    this.copy('nullTime.go', gormDir + 'nullTime.go');
    this.copy('gin.go', ginDir + 'gin.go');
    this.copy('routes.go', ginDir + 'routes.go');
    this.copy('objects.go', gormDir + 'objects.go');
    this.copy('gin_suite_test.go', ginDir + 'gin_suite_test.go');
    this.copy('user_test.go', ginDir + 'user_test.go');
    this.copy('Makefile', 'Makefile');

    var tmplContext = {
      myappName: this.myappName,
      myrepoUrl: this.myrepoUrl
    };

    this.template('_main.go', 'main.go', tmplContext);
    // this.template('_README.md', 'README.md');
  },

  // initialize git repo
  initializing: function () {
    this.spawnCommandSync('git', ['init', '--quiet']);
  },

  // start govendor directory
  install: function () {
    this.spawnCommand('go', ['get']);
    this.spawnCommand('go', ['get', '-u', 'github.com/kardianos/govendor']);
    this.spawnCommand('govendor', ['init']);
    // this.spawnCommand('govendor', ['add', '+external']);
    this.spawnCommand('make', ['install']);
  }

});
