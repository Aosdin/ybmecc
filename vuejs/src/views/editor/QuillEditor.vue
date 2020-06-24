<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl pt-0 mt-n3">
      <v-row>
        <app-card
          :heading="$t('message.quillEditor')"
          colClasses="col-xl-12 col-lg-12 col-md-12 col-sm-12 col-12"
        >
          <quill-editor v-model="content"
            ref="myQuillEditor"
            :options="editorOption"
            @blur="onEditorBlur($event)"
            @focus="onEditorFocus($event)"
            @ready="onEditorReady($event)"
          >
          </quill-editor>
        </app-card>
      </v-row>
		</v-container>
	</div>
</template>

<script>
export default {
  data: function() {
    return {
      content: "<h2>I am Example</h2>",
      editorOption: {
        // some quill options
      }
    };
  },
  methods: {
    onEditorBlur(editor) {
      console.log("editor blur!", editor);
    },
    onEditorFocus(editor) {
      console.log("editor focus!", editor);
    },
    onEditorReady(editor) {
      console.log("editor ready!", editor);
    },
    onEditorChange({ editor, html, text }) {
      console.log("editor change!", editor, html, text);
      this.content = html;
    }
  },
  computed: {
    editor() {
      return this.$refs.myQuillEditor.quill;
    }
  }
};
</script>
