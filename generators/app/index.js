'use strict';
var yeoman = require('yeoman-generator');
var mkdirp = require('mkdirp');

var chalk = require('chalk');
var yosay = require('yosay');

module.exports = yeoman.Base.extend({
  prompting: function () {
    // Have Yeoman greet the user.
    this.log(chalk.hex('#6fd6e3').bold(`
    ╭──────────────────────────╮
    │  Welcome to the terrific │
    │     generator-gin-api    │
    │        generator!        │
    ╰──────────────────────────╯
            ###
              ##
                #
                                    dhyysoo++++ooossyhd
                            Ndyo/:#######################:+sd
                  NN/:###ds/#######################/+////+/###+y#######
              Ny/:###/hy:#:+/::###::/+:#########+/:"  .:#  :+/##:oo######
             h:##//:+o:#:+:  +hdy.    #+:#####++"     s####" :+###:sN +###:
             ###d   /##:o   :#####      o:###//       h#s#N.  :/####+Ny####
            d###sN :###s"    "sd+d:     "o###s         ://"    s#####++###/
             y:##s/####s                "s###s"               "s######s:/
              Nhso#####//               +:###:o"             "o:######:
                N#######+/            "+/:+sss+o/"         ./+#########
                y########:+/:.    "#://#+N     N:+//:::::/+/###########/
                +############:/+++/:##ooodNNN dsoo+#####################
                /####################s/:::::::::::/s####################h
                :####################+o//+++s/++//oo####################s
                :######################o/  ##  ++/:#####################o
                /######################:+  ::  :/#######################+
                +#######################+//++//+########################+`));

    this.log(chalk.hex('#c0a98e').bold('\n' +
'                  +------------------------------------------------+\n' +
'                  | G o  G I N | p r o j e c t | g e n e r a t o r |\n' +
'                  +------------------------------------------------+\n' +
      '\n'));

    return this.prompt([{
      type: 'input',
      name: 'myAppPath',
      message: 'What is the root path for your project (ex: github.com/sezzle/generator-gin-api)',
      default: process.cwd().replace(process.env.GOPATH + "/src/", '')
    }]).then(function (answers) {
      this.log('app name', answers.myAppPath);
      this.myAppPath = answers.myAppPath;
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
        myAppPath : this.myAppPath,
    };

    this.template('_main.go', 'main.go', tmplContext);

  },

  // initialize git repo
  initializing: function () {
    console.log('Initializing git repository');
    this.spawnCommandSync('git', ['init', '--quiet']);
  },

  // start govendor directory
  install: function () {
    console.log('Running go get on dependencies...');
    this.spawnCommandSync('go', ['get']);

    console.log('Updating installation of govendor');
    this.spawnCommandSync('go', ['get', '-u', 'github.com/kardianos/govendor']);

    console.log('Initializing govendor');
    this.spawnCommandSync('govendor', ['init']);

    console.log('Vendoring external dependencies');
    this.spawnCommandSync('make', ['vendor-dependencies']);

    console.log('Vendoring Gomega (see https://github.com/onsi/gomega/issues/156)');
    this.spawnCommandSync('govendor', ['fetch', 'github.com/onsi/gomega']);
  }

});
