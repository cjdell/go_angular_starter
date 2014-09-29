var SignInController = ['$scope', '$state', '$stateParams', 'Authenticator',
  function($scope, $state, $stateParams, Authenticator) {
    'use strict';

    $scope.user = {
      email: 'a@b.c',
      password: 'password'
    };

    console.log($stateParams);

    $scope.signIn = function() {
      // if ($scope.signInForm.$pristine) {
      //   alert('Form is empty');
      //   return;
      // }

      if (!$scope.signInForm.$valid) {
        alert('Form is invalid');
        return;
      }

      Authenticator.signIn($scope.user.email, $scope.user.password).then(signedIn, failed);
    };

    function signedIn() {
      if ($stateParams.attemptedStateParams)
        $state.go($stateParams.attemptedStateName, JSON.parse(atob(decodeURI($stateParams.attemptedStateParams))));
      else
        $state.go('records.products.new');
    }

    function failed(error) {
      alert(error);
    }
  }
];

var RegisterController = ['$scope', '$state', '$stateParams', 'AuthApi',
  function($scope, $state, $stateParams, AuthApi) {
    'use strict';

    var user = $scope.user = {};

    $scope.register = function() {
      if ($scope.registerForm.$pristine) {
        alert('Form is empty');
        return;
      }

      if (!$scope.registerForm.$valid) {
        alert('Form is invalid');
        return;
      }

      AuthApi.Register(user.email, user.password).then(registerSuccessful, registerFailed);
    };

    function registerSuccessful() {
      $state.go('auth.sign-in');
    }

    function registerFailed(error) {
      alert(error);
    }

    $scope.$watch('registerForm.password.$viewValue', function() {
      var form = $scope.registerForm;

      if (form.$dirty) {
        form.password.$setValidity('length', form.password.$viewValue && form.password.$viewValue.length >= 3);
      }
    });
  }
];

module.exports = {
  SignInController: SignInController,
  RegisterController: RegisterController
};