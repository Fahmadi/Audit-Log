let auth = "";
window.onload = function () {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: "/swagger/apidocs.swagger.json",
    dom_id: "#swagger-ui",
    deepLinking: true,
    store: {
      auth: "",
    },
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    plugins: [SwaggerUIBundle.plugins.DownloadUrl],
    layout: "StandaloneLayout",
    requestInterceptor: (req) => {
      return req;
    },
    responseInterceptor: (res) => {
      if (res.url.includes("/v1/admin/users") && res.status == 200) {
        this.auth = JSON.parse(data).id;
      }
      return res;
    },
  });

  //</editor-fold>
};
