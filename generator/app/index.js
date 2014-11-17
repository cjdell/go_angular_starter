'use strict';
var util = require('util');
var path = require('path');
var yeoman = require('yeoman-generator');
var yosay = require('yosay')
var plural = require('plural');
var changeCase = require('change-case');
var esprima = require('esprima');
var escodegen = require('escodegen');

var GoAngularStarterGenerator = yeoman.generators.Base.extend({
  initializing: function () {
    this.pkg = require('../package.json');
  },

  prompting: function () {
    var done = this.async();

    // Have Yeoman greet the user.
    this.log(yosay(
      'Welcome to the Go Angular Starter entity generator!'
    ));

    var prompts = [{
      type: 'input',
      name: 'entityName',
      message: 'Specify entity type name (i.e SampleEntity):',
      default: 'SampleEntity'
    }];

    this.prompt(prompts, function (props) {
      this.entityNameSingularPascalCase = plural(props.entityName, 1);
      this.entityNamePluralPascalCase = plural(props.entityName, 2);

      this.entityNameSingularSnakeCase = changeCase.snakeCase(this.entityNameSingularPascalCase);
      this.entityNamePluralSnakeCase = changeCase.snakeCase(this.entityNamePluralPascalCase);

      this.entityNameSingularCamelCase = changeCase.camelCase(this.entityNameSingularPascalCase);
      this.entityNamePluralCamelCase = changeCase.camelCase(this.entityNamePluralPascalCase);

      done();
    }.bind(this));
  },

  writing: {
    app: function () {
      this.template('api.go', path.join('api', this.entityNameSingularSnakeCase + '.go'));
      this.template('service.go', path.join('services', this.entityNameSingularSnakeCase + '.go'));
      this.template('entity.go', path.join('model', 'entity', this.entityNameSingularSnakeCase + '.go'));
      this.template('persister.go', path.join('model', 'persister', this.entityNameSingularSnakeCase + '.go'));

      this.template('controller.js', path.join('admin', 'js', 'controllers', this.entityNameSingularSnakeCase + '.js'));
      this.template('api.js', path.join('admin', 'js', 'api', this.entityNameSingularSnakeCase + '.js'));

      this.template('index.html', path.join('admin', 'views', this.entityNamePluralSnakeCase, 'index.html'));
      this.template('view.html', path.join('admin', 'views', this.entityNamePluralSnakeCase, 'view.html'));

      // --------------------------------

      var serverPath = "server.go",
          serverSource = this.readFileAsString(serverPath);

      serverSource = serverSource.replace("// GENERATOR INJECT", "// GENERATOR INJECT\n\
\n\
"+this.entityNameSingularCamelCase+"Api := http.StripPrefix(\"/"+this.entityNamePluralSnakeCase+"\", api.New"+this.entityNameSingularPascalCase+"Api(db))\n\
\n\
apiRouter.Handle(\"/"+this.entityNamePluralSnakeCase+"\", "+this.entityNameSingularCamelCase+"Api)\n\
apiRouter.Handle(\"/"+this.entityNamePluralSnakeCase+"/{id:[0-9]+}\", "+this.entityNameSingularCamelCase+"Api)");

      this.write(serverPath, serverSource);

      // --------------------------------

      var appPath = "admin/js/app.js",
          appSource = this.readFileAsString(appPath);

      appSource = appSource.replace("// GENERATOR INJECT API", "// GENERATOR INJECT API\n\
app.factory('"+this.entityNameSingularPascalCase+"Api', require('./api/"+this.entityNameSingularSnakeCase+"'));");

      appSource = appSource.replace("// GENERATOR INJECT CTRL", "// GENERATOR INJECT CTRL\n\
\n\
var "+this.entityNameSingularPascalCase+"Controllers = require('./controllers/"+this.entityNameSingularSnakeCase+"');\n\
\n\
app.controller('"+this.entityNamePluralPascalCase+"Controller', "+this.entityNameSingularPascalCase+"Controllers."+this.entityNamePluralPascalCase+"Controller);\n\
app.controller('"+this.entityNameSingularPascalCase+"Controller', "+this.entityNameSingularPascalCase+"Controllers."+this.entityNameSingularPascalCase+"Controller);");

      this.write(appPath, appSource);

      // --------------------------------

      var recordsPath = "admin/views/layouts/records.html",
          recordsSource = this.readFileAsString(recordsPath);

      recordsSource = recordsSource.replace("<!-- GENERATOR INJECT -->", "<!-- GENERATOR INJECT -->\n\
<li ng-show=\"canView('"+this.entityNameSingularPascalCase+"')\">\n\
    <a ui-sref=\"records."+this.entityNamePluralSnakeCase+".new\">"+this.entityNamePluralPascalCase+"<span class=\"count\" ng-show=\"itemCounts."+this.entityNameSingularPascalCase+"\"> ({{itemCounts."+this.entityNameSingularPascalCase+"}})</span></a>\n\
</li>");

      this.write(recordsPath, recordsSource);

      // --------------------------------

      var routesPath = "admin/config/routes.json",
          routesSource = this.readFileAsString(routesPath);

      var routes = JSON.parse(routesSource);

      routes['records.' + this.entityNamePluralSnakeCase] = {
        "abstract": true,
        "allow": ["Admin"],
        "controller": this.entityNamePluralPascalCase + "Controller",
        "templateUrl": "views/" + this.entityNamePluralSnakeCase + "/index.html",
        "url": "/" + this.entityNamePluralSnakeCase
      };

      routes['records.' + this.entityNamePluralSnakeCase + '.new'] = {
        "allow": ["Admin"],
        "controller": this.entityNameSingularPascalCase + "Controller",
        "templateUrl": "views/" + this.entityNamePluralSnakeCase + "/view.html",
        "url": "/new"
      };

      routes['records.' + this.entityNamePluralSnakeCase + '.view'] = {
        "allow": ["Admin"],
        "controller": this.entityNameSingularPascalCase + "Controller",
        "templateUrl": "views/" + this.entityNamePluralSnakeCase + "/view.html",
        "url": "/{id:[0-9]+}"
      };

      routesSource = JSON.stringify(routes, null, 2);

      this.write(routesPath, routesSource);
    },

    projectfiles: function () {

    }
  },

  end: function () {
    this.installDependencies();
  }
});

module.exports = GoAngularStarterGenerator;
