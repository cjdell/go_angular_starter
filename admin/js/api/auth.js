module.exports = ['$http', 'Restangular',
  function($http, Restangular) {
    return {
      signIn: function(email, password) {
        return Restangular.all('auth/sign-in').post({
          Email: email,
          Password: password
        });
      },
      signUp: function(email, password) {
        return Restangular.all('auth/sign-up').post({
          Email: email,
          Password: password
        });
      }
    };
  }
];