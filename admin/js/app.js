/* APP
-------------------------------- */
require('angular');

var angular = window.angular;

var router = require('angular-ui-router'),
  ngDialog = require('./lib/ngDialog');

window._ = require('underscore'); // Restangular wants _ in global scope

require('angular-animate'); // Pollutes global scope, could do with improving
require('restangular');

var app = angular.module('go-angular-starter', ['ngLocale', 'ngAnimate', 'restangular', router, ngDialog]);

/* SERVICES
-------------------------------- */
app.factory('Authenticator', require('./services/authenticator'));
app.factory('Uploader', require('./services/uploader'));
app.factory('Utility', require('./services/utility'));

/* API
-------------------------------- */
app.factory('AuthApi', require('./api/auth'));
app.factory('UserApi', require('./api/user'));
app.factory('ProductApi', require('./api/product'));
app.factory('CategoryApi', require('./api/category'));

/* DIRECTIVES
-------------------------------- */
app.directive('match', require('./directives/match'));
app.directive('fileUpload', require('./directives/file_upload'));
app.directive('input', require('./directives/blur_focus'));
app.directive('select', require('./directives/blur_focus'));
app.directive('categoryField', require('./directives/category_field'));
app.directive('tinyMce', require('./directives/tiny_mce'));
app.directive('parentWidth', require('./directives/parent_width'));
app.directive('scrollFixed', require('./directives/scroll_fixed'));
app.directive('scrollContainer', require('./directives/scroll_container'));

/* CONTROLLERS
-------------------------------- */
var AuthControllers = require('./controllers/auth'),
  UserControllers = require('./controllers/user'),
  ProductControllers = require('./controllers/product'),
  CategoryControllers = require('./controllers/category');

app.controller('SignInController', AuthControllers.SignInController);
app.controller('RegisterController', AuthControllers.RegisterController);

app.controller('UsersController', UserControllers.UsersController);
app.controller('UserController', UserControllers.UserController);

app.controller('ProductsController', ProductControllers.ProductsController);
app.controller('ProductController', ProductControllers.ProductController);

app.controller('CategoriesController', CategoryControllers.CategoriesController);
app.controller('CategoryController', CategoryControllers.CategoryController);
app.controller('SelectCategoryController', CategoryControllers.SelectCategoryController);

app.run(require('./run/auth'));
app.run(require('./run/item_count'));
app.run(require('./run/permissions'));
app.run(require('./run/promise'));
app.run(require('./run/state_class'));

app.config(require('./config/http'));
app.config(require('./config/restangular'));
app.config(require('./config/router'));

// Manual Angular bootstrap call, less magic
angular.bootstrap(window.document.body, ['go-angular-starter']);