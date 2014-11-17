var angular = require('angular'),
  _ = require('underscore'),
  common = require('./common');

var <%= entityNamePluralPascalCase %>Controller = ['$injector', '$scope', '$state', '$stateParams', '<%= entityNameSingularPascalCase %>Api',
  function($injector, $scope, $state, $stateParams, <%= entityNameSingularPascalCase %>Api) {
    var ctrl = this;

    // Inherit shared functionality from the ListController
    $injector.invoke(common.ListController, ctrl, {
      $scope: $scope,
      Api: <%= entityNameSingularPascalCase %>Api
    });

    // Bootstrap the controller
    ctrl.init();
  }
];

var <%= entityNameSingularPascalCase %>Controller = ['$injector', '$scope', '$state', '$stateParams', '<%= entityNameSingularPascalCase %>Api',
  function($injector, $scope, $state, $stateParams, <%= entityNameSingularPascalCase %>Api) {
    var ctrl = this;

    // Inherit shared functionality from the ItemController
    $injector.invoke(common.ItemController, ctrl, {
      $scope: $scope,
      $state: $state,
      $stateParams: $stateParams,
      Api: <%= entityNameSingularPascalCase %>Api
    });

    // Bootstrap the controller
    ctrl.init();
  }
];

module.exports = {
  <%= entityNamePluralPascalCase %>Controller: <%= entityNamePluralPascalCase %>Controller,
  <%= entityNameSingularPascalCase %>Controller: <%= entityNameSingularPascalCase %>Controller
};