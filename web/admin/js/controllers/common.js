var _ = require('underscore'),
  Validate = require('validate-arguments');

var ListController = ['$scope', 'Api',
  function($scope, Api) {
    var args = Validate.validate(arguments, ['object', {
      EntityNames: 'object',
      GetAll: 'function'
    }]);

    if (!args.isValid()) {
      throw args.errorString();
    }

    var ctrl = this; // Only time `this` is used. Save us from confusion...

    ctrl.load = function() {
      Api.GetAll().then(ctrl.loaded, ctrl.failed);
    };

    ctrl.loaded = function(reply) {
      $scope.records = reply[Api.EntityNames.PluralPascalCase];

      console.table($scope.records);

      $scope.$emit('itemCountDiscovered', Api.EntityNames.SingularPascalCase, $scope.records.length);
    };

    ctrl.failed = function(err) {
      console.error(err);
    };

    $scope.$on('item-saved', this.load);

    $scope.$on('item-deleted', this.load);

    // Load the list
    ctrl.load();
  }
];

var ItemController = ['$scope', '$state', '$stateParams', 'Api',
  function($scope, $state, $stateParams, Api) {
    var args = Validate.validate(arguments, ['object', 'object', 'object', {
      EntityNames: 'object',
      GetOne: 'function',
      Update: 'function',
      Insert: 'function',
      Delete: 'function'
    }]);

    if (!args.isValid()) {
      throw args.errorString();
    }

    var ctrl = this; // Only time `this` is used. Save us from confusion...

    ctrl.load = function(id) {
      Api.GetOne(id).then(ctrl.loaded, ctrl.failed);
    };

    // These functions have been exposed to the view via the scope
    ctrl.new = function() {
      ctrl.goToRecord();
    };

    ctrl.validate = function() {
      return true;
    };

    ctrl.save = function() {
      if (ctrl.validate()) {
        if ($scope.record.Id !== undefined) {
          ctrl.update();
        } else {
          ctrl.insert();
        }
      }
    };

    ctrl.update = function() {
      Api.Update($scope.record).then(ctrl.saved, ctrl.failed);
    };

    ctrl.insert = function() {
      Api.Insert($scope.record).then(ctrl.saved, ctrl.failed);
    };

    ctrl.delete = function() {
      Api.Delete($scope.record.Id).then(ctrl.deleted, ctrl.failed);
    };

    // Callbacks are declared privately
    ctrl.loaded = function(reply) {
      $scope.record = reply[Api.EntityNames.SingularPascalCase];

      console.log('Loaded Record:', $scope.record);
    };

    ctrl.failed = function(err) {
      console.error(err);
    };

    ctrl.saved = function(reply) {
      $scope.record = reply[Api.EntityNames.SingularPascalCase];

      console.log('Saved Record:', $scope.record);

      // Broadcast saved event
      $scope.$emit('item-saved', $scope.record);

      ctrl.goToRecord($scope.record.Id);
    };

    ctrl.deleted = function() {
      console.log('Deleted Record:', $scope.record);

      // Broadcast deleted event
      $scope.$emit('item-deleted', $scope.record);

      // Go to a blank record
      ctrl.goToRecord();
    };

    ctrl.goToRecord = function(id) {
      if (id) {
        $state.go('records.' + Api.EntityNames.PluralSnakeCase + '.view', {
          id: id
        });
      } else {
        $state.go('records.' + Api.EntityNames.PluralSnakeCase + '.new');
      }
    };

    // These are the methods we choose to expose to the view via the $scope
    _.extend($scope, _.pick(this, ['new', 'save', 'delete']));

    // If there is an ID, load the record from the API, otherwise, start with a new instance
    if ($stateParams.id !== undefined) {
      ctrl.load(parseInt($stateParams.id));
    }
  }
];

module.exports = {
  ListController: ListController,
  ItemController: ItemController
};