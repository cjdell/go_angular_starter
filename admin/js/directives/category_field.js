var _ = require('underscore');

module.exports = ['ngDialog',
  function(ngDialog) {
    return {
      restrict: 'A',
      scope: {
        categoryField: '=',
        categoryMode: '@'
      },
      templateUrl: 'views/directives/category_field.html',
      controller: function($scope) {
        var dialog = null;

        if (['one', 'many'].indexOf($scope.categoryMode) === -1) {
          console.error('Invalid categoryMode:', $scope.categoryMode);
          return;
        }

        // Make sure we have the correct data type
        if ($scope.categoryMode === 'many' && !($scope.categoryMode instanceof Array)) {
          $scope.categoryField = [];
        }

        $scope.getSelectedText = function() {
          if ($scope.categoryMode === 'one') {
            return $scope.categoryField ? $scope.categoryField.FqName : '[None]';
          } else {
            return $scope.categoryField.length;
          }
        };

        $scope.selectCategory = function() {
          dialog = ngDialog.open({
            template: 'views/dialogs/select_category.html',
            className: 'ngdialog-theme-plain select-category-dialog',
            scope: $scope,
            controller: 'SelectCategoryController'
          });

          dialog.closePromise.then(gotSelection);
        };

        $scope.clearCategory = function() {
          if ($scope.categoryMode === 'one') {
            $scope.categoryField = null;
          } else {
            $scope.categoryField = [];
          }
        };

        $scope.removeCategory = function(category) {
          $scope.categoryField.splice($scope.categoryField.indexOf(category), 1);
        };

        function gotSelection(data) {
          if (typeof data.value === 'object') {
            if ($scope.categoryMode === 'one') {
              $scope.categoryField = data.value;
            } else {
              $scope.categoryField.push(data.value);
            }
          }
        }
      },
      link: function(scope, element, attr, ctrl) {
        element.addClass('category-field');
      }
    };
  }
];