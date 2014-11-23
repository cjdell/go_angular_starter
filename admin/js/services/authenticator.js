module.exports = ['$http', 'AuthApi',
  function($http, AuthApi) {
    var user = null;

    return {
      isAuthenticated: function() {
        return user !== null;
      },
      signIn: function(email, password) {
        return AuthApi.signIn(email, password).then(function(reply) {
          $http.defaults.headers.common['API-Key'] = reply.ApiKey;

          user = {
            ApiKey: reply.ApiKey,
            Email: reply.Email,
            Name: reply.Name,
            Type: reply.Type
          };

          return user;
        });
      },
      signOut: function() {
        user = null;
        delete $http.defaults.headers.common['API-Key'];
      },
      getUser: function() {
        return user;
      }
    };
  }
];