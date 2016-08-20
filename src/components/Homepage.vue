<template>
  <main>
    <div class="row" id="login">
      <div class="col s12 m10 offset-m1 l8 offset-l2">
        <div class="card blue-grey darken-1">
          <div class="card-image">
            <img src="http://2.bp.blogspot.com/-7ywyKE1iKaA/Tyfbcob7R3I/AAAAAAAAIwQ/h83RQebeEJ8/s1600/2012-01-29_22.15.39.png">
            <span class="card-title">Try it!</span>
          </div>
          <div class="card-content white-text">
            <p>想體驗一下嗎?</p>
          </div>
          <div class="card-action">
            <a class="waves-effect waves-light blue darken-2 btn" onclick="checkLoginState();"><i class="fa fa-facebook-official"></i>&nbsp;&nbsp;facebook login</a>
            <!--
                          <div class="fb-login-button" data-max-rows="3" data-size="large" data-show-faces="true" data-auto-logout-link="true" onclick="checkLoginState();"></div>
                          -->
          </div>
        </div>
      </div>
    </div>
  </main>

</template>
<style src="../css/style.css"/>
<script>
import '../js/fb_api'
import '../js/myjs'

</script>
