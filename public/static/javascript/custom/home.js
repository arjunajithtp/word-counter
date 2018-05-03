myApp.controller('AppCtrl', ['$scope', function($scope) {

    $scope.keys = [];
    $scope.gridData = {};

    $scope.getWordCount = function() {

        $.ajax({
            url: '/word-counter/home',
            type: 'POST',
            dataType: 'json',
            data : "&url=" + $scope.url,
            success : function(data) {
                $scope.$apply(function(){
                    $.extend(true,$scope.gridData,data);
                    $scope.keys = Object.keys($scope.gridData);
                });
            }
        });
    };


}]);