<template>
  <main>
    <div class="row" >
      <div class="col s12 m10 offset-m1 l8 offset-l2">
        <div class="card blue-grey darken-1">
          <div v-if="statuss ==1">
          <div class="card-image">
            <img src="http://2.bp.blogspot.com/-7ywyKE1iKaA/Tyfbcob7R3I/AAAAAAAAIwQ/h83RQebeEJ8/s1600/2012-01-29_22.15.39.png">
            <span class="card-title">Try it!</span>
          </div>

          <div class="card-content white-text">
            <p>想體驗一下嗎?</p>
          </div>
            </div>
          <div v-else>

                  <div class="card-content white-text">
                    <span class="card-title\" id="{{ username }}"></span>
                    <p>Profile Information goes here</p>
                  </div>

            </div>
          <div class="card-action">
            <div v-if="statuss == 1">
            <a class="waves-effect waves-light blue darken-2 btn" v-on:click="checkLoginState"><i class="fa fa-facebook-official"></i>&nbsp;&nbsp;facebook login</a>
            </div>
            <div v-else>
              <a class="waves-effect waves-light blue darken-2 btn" v-on:click="IWantToLogout"><i class="fa fa-facebook-official"></i>&nbsp;&nbsp;facebook logout</a>
            </div>
        </div>
      </div>
    </div>
  </main>

</template>

<script>


  export default {
    data ()
    {
      return {
      statuss : 1 ,
        username : ''
    }
  },
  ready: function () {
      var self = this;
      window.fbAsyncInit = function () {
        FB.init({
          appId: '1014922658562861',
          cookie: true,  // enable cookies to allow the server to access
          // the session
          xfbml: true,  // parse social plugins on this page
          version: 'v2.7' // use graph api version 2.5
        });
        FB.getLoginStatus(function (response) {
          self.statusChangeCallback(response);
        });
      };
      (function (d, s, id) {
        var js, fjs = d.getElementsByTagName(s)[0];
        if (d.getElementById(id)) return;
        js = d.createElement(s);
        js.id = id;
        js.src = "//connect.facebook.net/zh_TW/sdk.js#xfbml=1&version=v2.7&appId=1014922658562861";
        fjs.parentNode.insertBefore(js, fjs);
      }(document, 'script', 'facebook-jssdk'));
    },
    methods: {
      IWantToLogout :function (){
        var self=this
        FB.logout(function (response) {
          self.statuss = 1
          $("ul li:nth-child(3)")[0].innerHTML = "<li><a href=\"#\">Login</a></li>";
        });
      },
      statusChangeCallback : function(response) {
        //alert('status')
        var self = this
        console.log(self.statuss)
        if (response.status === 'connected') {
         // $("#login").hide();
          self.statuss = 2
          FB.api('/me', function (response) {
            //$('#UserName')[0].innerHTML = response.name;
            alert(response.name);
            self.username=response.name
            $("ul li:nth-child(3)")[0].innerHTML = "<li><a href=\"#\">" + response.name + "</a></li>";
          });
          var posts = this.getPost();
          /*
           $.post(receive.go, {
           name: UserName;
           text: posts;
           });
           */
        }
      },
      getPost: function()
      {
        FB.api('/me/posts', function (response) {
          //console.log(response.data);
          var data = response.data
          var ans = [];
          for (var i = 0; i < data.length; i++) {
            var tmp = {};
            var str = "";
            if ('story' in data[i])str += data[i].story;
            if ('message' in data[i])str += data[i].message;
            tmp.message = str;
            tmp.created_time = data[i].created_time;
            ans.push(tmp);
          }
         // console.log(ans);
          return ans;
        });
      },

      checkLoginState: function () {
        var self = this
        //alert('hi')
        FB.getLoginStatus(function (response) {
          if (response.status === 'connected') {
            this.statusChangeCallback(response);
          }
          else {
            FB.login(function (response) {
            //  console.log(response)
              self.statusChangeCallback(response);
            }, {scope: 'email,public_profile,user_posts', return_scopes: true});
          }
        });
      }
    }
  }
</script>
