var angular = require('angular'),
  _ = require('underscore'),
  common = require('./common');

var CategoriesController = ['$injector', '$scope', '$state', '$stateParams', 'CategoryApi',
  function($injector, $scope, $state, $stateParams, CategoryApi) {
    var ctrl = this;

    // Inherit shared functionality from the ListController
    $injector.invoke(common.ListController, ctrl, {
      $scope: $scope,
      Api: CategoryApi
    });

    // Bootstrap the controller
    ctrl.init();
  }
];

var CategoryController = ['$injector', '$scope', '$state', '$stateParams', 'CategoryApi',
  function($injector, $scope, $state, $stateParams, CategoryApi) {
    var ctrl = this;

    // Inherit shared functionality from the ItemController
    $injector.invoke(common.ItemController, ctrl, {
      $scope: $scope,
      $state: $state,
      $stateParams: $stateParams,
      Api: CategoryApi
    });

    $scope.$watch('record.Parent', function() {
      if ($scope.record) {
        if ($scope.record.Parent) {
          $scope.record.ParentId = $scope.record.Parent.Id;
        } else {
          $scope.record.ParentId = 0;
        }
      }
    });

    // Bootstrap the controller
    ctrl.init();
  }
];

var SelectCategoryController = ['$scope', 'CategoryApi',
  function($scope, CategoryApi) {
    $scope.name = 'Select category';

    $scope.levels = [];
    $scope.levelIndex = -1;

    $scope.nextLevel = function(parent) {
      CategoryApi.getAll({
        parent_id: parent ? parent.Id : 0
      }).then(function(reply) {
        $scope.levels.push({
          parent: parent,
          categories: reply
        });
        $scope.levelIndex++;
      });
    };

    $scope.select = function(category) {
      $scope.closeThisDialog(category);
    };

    $scope.nextLevel(null);
  }
];

module.exports = {
  CategoriesController: CategoriesController,
  CategoryController: CategoryController,
  SelectCategoryController: SelectCategoryController
};