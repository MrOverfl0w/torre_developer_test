<template>
  <div class="content">
    <header>
      <b-container fluid>
        <div class="logo">
            <img id="logoTorre" alt="Torre"  src="https://ci3.googleusercontent.com/proxy/I_RyBRVKmDfXYBzgHoYHsCmDGyPgn2bC9u6wPa0SypxG5acS-10Sxnk85Ma4qqBEoRIITMM0YSm5OJQfonvVDt3DlStIiDWL2uXRna_rIG_NbedIsQ-W=s0-d-e1-ft#https://s3-us-west-2.amazonaws.com/torre-media/static/torre_dark.png">
        </div>
        <div class="buscar">
            <b-form-input v-model="skill" placeholder="Search by skill" name="skill"></b-form-input>
            <b-button v-on:click="postreq()">Search Mentor</b-button>
        </div>
        <div class="signin"><a href="https://accounts.torre.co/email/?next=/openid/authorize%3Fscope%3Dopenid%2Bprofile%2Bemail%26response_type%3Dcode%26redirect_uri%3Dhttps%253A%252F%252Ftorre.co%252Fcallback%253Fclient_name%253Dstarrgate%26state%3DKz7-TYl8rMFi5xhMNfnBlZOJQXf6dOl2A_ZqIWRJZ4E%26intent%3Dhome%253Asign-in%26client_id%3D541493">Go to Torre</a></div>
      </b-container>
    </header>
    <div class="content-body">
      <div class="banner">
        <div class="texto-banner">
          <h1>Torre: Mentors</h1>
        </div>
      </div>
      <div class="resContainer">
        <ul class="results">
          <li class="result" v-for="person in people" :key="person.id">
            <a :href="person.url" target="_blank">
              <div class="user-image"><img :src="person.picture" alt=""></div>
              <div class="user-content">
                <div class="name"><span>{{ person.name }}</span></div>
                <div class="price"><span>{{ person.price }}</span></div>
                <div class="remote" v-if="person.remote"><span>Remote</span></div>
              </div>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>

import axios from 'axios';

export default {
  name: 'MentorsFront',

  data: function(){
    return {
      skill: "",
      people: ""
    }
  },

  methods: {
    postreq: function(){
      var data = {"skill": this.skill}
      console.log(data);
      axios({ method: "POST", url: "http://127.0.0.1:8090/", data: data, headers: {"content-type": "text/plain" } }).then( result => {
        this.people = result.data["results"]
        console.log(result.data);
      }).catch( error => {
        console.log(error);
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.content{
  letter-spacing: 0.4px;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  margin: 10px;
}
/* Header */
header{
  position: sticky;
}
header .container-fluid{
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 1px 3px rgb(0 0 0 / 20%), 0 1px 1px rgb(0 0 0 / 14%), 0 2px 1px -1px rgb(0 0 0 / 12%);
  padding: 13px 3%;
}
.logo #logoTorre{
  max-height: 17px;
  width: auto;
}
.buscar{
  display: flex;
}
.buscar input{
  width: 250px;
  border: 1px solid #FFFFFF60;
  border-radius: 25px;
  background-color: #FFFFFF30;
  font-size: 0.8em;
  color: #FFF;
}
.buscar button{
  font-size: 0.7em;
  background: transparent;
  border: none;
  padding: 3px 15px;
}
.signin a{
  font-size: 1em;
  font-weight: 600;
  color: #cddc39;
  text-decoration: none;
}

/* Content */
.banner{
  min-height: 50vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-image: url("https://res.cloudinary.com/torre-technologies-co/image/upload/f_auto,q_auto:good/v1614947275/origin/explorer/csei905chcsrxycc2nwv.png");
  background-position: center center;
  background-repeat: repeat-x;
  padding: 50px 3%;
}
.texto-banner{
  font-size: 1.6em;
  font-weight: bold;
  letter-spacing: 1px;
  color: 27292d;
}

.resContainer{
  margin: 40px 0;
  padding: 30px 3%;
}
.results{
  display: grid;
  grid-template-columns: repeat(3,minmax(0,1fr));
  grid-row-gap: 20px;
  grid-column-gap: 20px;
}
.results .result{
  border-radius: 25px;
  box-shadow: 0px 0px 17px 0px rgb(0 0 0 / 20%);
}
.results .result > a{
  padding: 20px;
  display: grid;
  grid-template-columns: 130px auto;
  text-decoration: none;
}
.result .user-image img{
  width: 100%;
}
.result .user-content{
  padding: 10px;
  color: #FFF;
  align-self: center;
}
.user-content .name{
  font-size: 1.1em;
}
.user-content .price{
  margin-top: 5px;
  font-size: 0.9em;
}
.user-content .remote{
  margin-top: 10px;
  font-size: 0.7em;
}
</style>
