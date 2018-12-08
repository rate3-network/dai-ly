<template>
  <v-app id="inspire">
    <v-navigation-drawer
      fixed
      v-model="drawer"
      app
    >
      <v-list dense>
        <v-list-tile @click="">
          <v-list-tile-action>
            <v-icon>home</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Home</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile @click="">
          <v-list-tile-action>
            <v-icon>contact_mail</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Contact</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar color="indigo" dark fixed app>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>Delegate Status</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <v-container fluid fill-height>
        <v-layout
          justify-center
          align-center
        >
          <v-card>
            <v-img
              src="data:image/png;base64,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"
              aspect-ratio="1"
              s
            ></v-img>

            <v-card-title primary-title>
              <div>
                <h3 class="headline mb-0">Kangaroo Valley Safari</h3>
                <div>Located two hours south of Sydney in the <br>Southern Highlands of New South Wales, ...</div>
              </div>
            </v-card-title>

            <v-card-actions>
              <v-btn flat color="orange">Share</v-btn>
              <v-btn flat color="orange">Explore</v-btn>
            </v-card-actions>
          </v-card>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Delegate',
  data: () => ({
    drawer: null
  }),
  mounted () {
    axios
      .get('https://api.coindesk.com/v1/bpi/currentprice.json')
      .then(response => {
        console.log(response)
        this.info = response
      })
  },
  props: {
    source: String
  }
}
</script>

<style>

</style>
