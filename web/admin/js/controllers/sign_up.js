var BusinessSignUpController = ['$scope', '$state', '$stateParams', 'BusinessWizard', 'SignUpApi',
  function($scope, $state, $stateParams, BusinessWizard, SignUpApi) {
    $scope.stageIndex = parseInt($stateParams.stageIndex);
    $scope.stageCount = 3;

    $scope.signUpData = BusinessWizard.signUpData;

    // Otherwise this will be blank
    $scope.password_confirm = BusinessWizard.signUpData.User.Password;

    $scope.goToStage = function(toStageIndex) {
      // var form = $scope['stageForm' + $scope.stageIndex];
      var form = $scope.stageForm;
      var goingForwards = toStageIndex > $scope.stageIndex;

      if (goingForwards && !form.$valid) {
        alert('Please complete the form before continuing');
        return;
      }

      // BusinessWizard.addStageData($scope.stageData);

      $scope.$parent.goToStage(toStageIndex);
    };

    $scope.finish = function() {
      SignUpApi.BusinessSignUp($scope.signUpData).then(function(reply) {
        $state.go('auth.sign-in');
      }, function(error) {
        alert(error);
      });
    };
  }
];

module.exports = {
  BusinessSignUpController: BusinessSignUpController
};