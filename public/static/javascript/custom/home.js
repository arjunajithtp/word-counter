myApp.controller('AppCtrl', ['$scope', function($scope) {

    $scope.getWordCount = function() {
        $.ajax({
            url: '/word-counter/home',
            type: 'POST',
            dataType: 'json',
            data : "&url=" + $scope.url,
            success : function(data) {
                console.log(data)
                $scope.$apply(function(){
                    $scope.words = data;
                    console.log($scope.words)
                });
            }
        });
    };


}]);