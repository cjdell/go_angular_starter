module.exports = ['$http', 'AuthApi',
  function($http, AuthApi) {
    return {
      isAuthenticated: function() {
        return !!$http.defaults.headers.common['API-Key'];
      },
      signIn: function(email, password) {
        return AuthApi.SignIn(email, password).then(function(reply) {
          $http.defaults.headers.common['API-Key'] = reply.ApiKey;

          return null;
        });
      },
      signOut: function() {
        delete $http.defaults.headers.common['API-Key'];
      }
    };
  }
];