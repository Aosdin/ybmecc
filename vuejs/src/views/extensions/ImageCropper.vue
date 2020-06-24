<template>
  <div class="img-cropper-wrap">
    <page-title-bar></page-title-bar>
    <v-container fluid class="pt-0 mt-n3">
      <v-row>
        <app-card
          colClasses="col-xl-12 col-lg-12 col-md-12 col-sm-12 col-12"
        >
          <v-row>
            <app-card
              colClasses="col-xl-6 col-lg-6 col-md-6 col-sm-6 col-12"
              customClasses="img-crop"
            >
              <croppa v-model="croppaImg" canvas-color="transparent"
                :placeholder="'Choose an image'"
                :accept="'image/*'"
                :show-remove-button="true"
                :remove-button-color="'red'"
                :initial-position="'center'"
              >
                <img slot="placeholder" src="public/static/img/bg-1.jpg" />
              </croppa>
              <div class="d-custom-flex justify-space-between py-3 format-btn-wrap">
                <v-btn @click="rotate()" class="primary">Rotate</v-btn>
                <v-btn @click="flipX()" class="primary">Flip-X</v-btn>
                <v-btn @click="flipY()" class="primary">Flip-Y</v-btn>
              </div>
              <v-btn class="v-btn success w-100 ma-0" @click="generateImage">Generate</v-btn>
            </app-card>
            <app-card
              v-if="imgUrl"
              colClasses="col-xl-6 col-lg-6 col-md-6 col-sm-6 col-12"
              customClasses="img-crop"
              contentCustomClass="d-custom-flex justify-center align-center h-100"
            >
              <img height="300" :src="imgUrl" class="img-crop img-responsive" >
            </app-card>
          </v-row>
        </app-card> 
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  data(){
    return {
      croppaImg: null,
      imgUrl: ''
    }
  },
  methods: {
    generateImage () {
    let url = this.croppaImg.generateDataUrl()
      if (!url) {
        alert('no image')
        return
      }
      this.imgUrl = url
    },
    rotate () {
      this.croppaImg.rotate()
    },
    flipX () {
      this.croppaImg.flipX()
    },
    flipY () {
      this.croppaImg.flipY()
    }
  }
}
</script>

