module.exports = ['$http', '$q',
  function($http, $q) {
    return {
      SignIn: function(email, password) {
        return $http.jsonrpc('/auth', 'AuthApi.SignIn', [{
          Email: email,
          Password: password
        }]);
      },
      Register: function(email, password) {
        return $http.jsonrpc('/auth', 'AuthApi.Register', [{
          Email: email,
          Password: password
        }]);
      }
    };
  }
];