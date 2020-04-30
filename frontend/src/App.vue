<template>
  <div id="app">
    <nav id="nav">
      <img src="./assets/logo_3.svg" alt="Covid and Colleges" @click="scrollTop()" />
    </nav>
    <div class="container">
      <div class="content">
        <header>
          <div class="info">
            <h1>Cache Performance Analyser</h1>
            <h2>A way to analyse the cache performance for the program you write for embedded systems</h2>
            
            <div id="main">
              <input type="file" @change="onFileChange" id="inputBox">
              <button v-on:click="makeRequest">
                Upload & Compile ⚙️
              </button>
            </div>
          </div>
          <div class="illustration">
            <img src="./assets/abc.png" alt="Fighting Corona Virus" />
          </div>
        </header>

        <div v-if="!loading" class="title">
            Hardware Performance Counter Values
          </div>
        <div class="collegeList">
          <div class="empty"></div>
          <div v-if="loading">
            <!-- <div class="dot1"></div>
            <div class="dot2"></div> -->
          </div>
          
          <div v-else v-for="(value, name) in this.programData" :key="name" id="tamjham">
              <div class="text">
                <span class="name">{{name}}</span> &ensp; -> &ensp; <span class="value">{{value}}</span>
              </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import "reset-css";
import axios from "axios";
window.document.body.onscroll = function() {
  if (
    document.body.scrollTop > 240 ||
    document.documentElement.scrollTop > 240
  ) {
    document.getElementById("nav").style.top = "0";
  } else {
    document.getElementById("nav").style.top = "-67px";
  }
};
export default {
  name: "app",
  data() {
    return {
      programData: undefined,
      loading: true,
      file: undefined
    };
  },
  methods: {
    makeRequest: function() {
      console.log("Ha")
      let vm = this;
      let formData = new FormData();
      formData.append('fileupload', vm.file);
      let headers = {'Content-Type': 'multipart/form-data'}
      console.log(formData, headers)
      axios
        .post("http://localhost:3000/uploadFile", formData, headers)
        .then(function(response) {
          vm.loading = false;
          vm.programData = response.data;
        })
        .catch(function(error) {
          // handle error
          console.log(error);
        });
    },
    onFileChange: function(e) {
      this.file = e.target.files[0] || e.dataTransfer.files;
      this.programData = undefined;
      this.loading = false;
    },
    scrollTop: function() {
      const c = document.documentElement.scrollTop || document.body.scrollTop;
      if (c > 0) {
        window.requestAnimationFrame(this.scrollTop);
        window.scrollTo(0, c - c / 10);
      }
    }
  }
};
</script>

<style lang="scss">

@media screen and (min-width: 1275px) {
  .empty-state {
    grid-column: 1 / span 2;
    margin-bottom: 4em;
    h2 {
      font-size: 20px;
      line-height: 26px;
      margin: 0.5em 0 0.3em 0;
    }
    h3 {
      font-size: 14px;
      line-height: 20px;
    }
  }
  nav {
    padding: 1em 8em;
    img {
      height: 24px;
    }
    button {
      margin-right: 24em;
      font-size: 12px;
    }
  }
  .container {
    padding: 2.5em;
    footer {
      .disc {
        font-size: 16px;
        line-height: 22px;
      }
      .showbazi {
        font-size: 14px;
        line-height: 20px;
      }
    }
    .content {
      padding: 4em;
      .disc {
        margin-top: 4em;
      }
      header {
        flex-direction: row;
        align-items: center;
        .info {
          max-width: 40%;
          margin-right: 120px;
          h1 {
            font-size: 1.75em;
          }
          h2 {
            font-size: 1.5em;
          }
          button {
            font-size: 1em;
          }
          p {
            margin-top: 3em;
          }
        }
      }
      .collegeList {
        grid-gap: 3rem;
        margin-top: 2em;
        grid-template-columns: repeat(auto-fit, minmax(500px, 400px));
        padding: 0 4em !important;
        justify-content: center;
        .collegeObj {
          padding: 1.5em 3em;
          h3 {
            font-size: 1.25em;
            line-height: 1.5em;
          }
          .collegeDetails {
            flex-direction: row;
            .row {
              display: flex;
              flex-direction: column;
            }
            .r1 {
              margin-right: 4em;
            }
            .section:first-child {
              margin: 0.5em 0 2em 0;
            }
            .section {
              h4 {
                margin-bottom: 0.5em;
                font-size: 0.75em;
              }
              p {
                padding: 0.5em 1em;
                margin-top: 0.5em;
                font-size: 0.875em;
                line-height: 1.5em;
              }
            }
          }
        }
      }
    }
  }
}
nav {
  display: flex;
  flex-direction: row;
  align-items: center;
  background: #ffffff;
  box-shadow: 0px 4px 8px rgba(68, 68, 68, 0.05);
  position: fixed;
  justify-content: space-between;
  width: 100%;
  z-index: 999;
  top: -67px;
  transition-duration: 0.3s;
  transition-timing-function: ease-in-out;
  transition-property: all;
  img {
    cursor: pointer;
  }
}
.empty-state {
  text-align: center;
  img {
    height: 72px;
  }
  h2 {
    font-family: "IBM Plex Sans", sans-serif;
    font-weight: 600;
    color: #462a4f;
  }
  h3 {
    font-family: "IBM Plex Sans", sans-serif;
    font-style: normal;
    font-weight: normal;
    color: rgba(70, 42, 79, 0.5);
    a {
      text-decoration-line: underline;
      text-decoration-style: dotted;
      transition-duration: 0.3s;
      transition-timing-function: ease-in-out;
      transition-property: all;
      color: #6d72c5;
    }
  }
}
header,
nav {
  button {
    font-family: "IBM Plex Mono", monospace;
    font-weight: 600;
    padding: 8px 12px;
    background: #7bbf75;
    border: 1px solid #6aa565;
    box-sizing: border-box;
    border-radius: 2px;
    cursor: pointer;
    align-self: flex-start;
  
    color: #ffffff;
    text-decoration: none;
  
  }
}
.search {
  input {
    padding: 0.5em;
    font-family: "IBM Plex Sans", sans-serif;
    color: #777;
    border: 1px solid #ddd;
    border-radius: 2px;
    font-size: 0.875em;
    transition-duration: 0.3s;
    transition-timing-function: ease-in-out;
    transition-property: all;
    background-image: url(./assets/search.svg);
    background-repeat: no-repeat;
    background-position: 8px 8px;
    background-size: 16px 16px;
    text-indent: 24px;
    &:focus {
      outline: none;
      border: 1px solid #6d72c5;
      background-color: #fff;
    }
  }
}
::-webkit-input-placeholder {
  /* Edge */
  color: #aaa;
}
.container {
  background-color: #f6f6f6;
  .content {
    max-width: 1192px;
    height: 1200px;
    margin: 0 auto;
    background-color: #ffffff;
    box-shadow: 0px 6px 12px rgba(43, 143, 143, 0.1);
    border-radius: 2px;
    footer {
      text-align: center;
      font-family: "IBM Plex Sans", sans-serif;
      a {
        text-decoration-line: underline;
        text-decoration-style: dotted;
        transition-duration: 0.3s;
        transition-timing-function: ease-in-out;
        transition-property: all;
      }
      .disc {
        color: #777;
        margin-bottom: 1em;
        a {
          color: #6d72c5;
          &:hover {
            color: #484fb7;
          }
        }
      }
      .showbazi {
        color: #444;
        a {
          color: #cc3660;
        }
      }
    }

    header {
      display: flex;
      justify-content: center;
      .info {
        display: flex;
        flex-direction: column;
        h1 {
          font-family: "IBM Plex Sans", sans-serif;
          font-weight: 600;
          line-height: 1.5em;
          color: #6d72c5;
        }
        h2 {
          font-family: "IBM Plex Sans", sans-serif;
          font-weight: 400;
          line-height: 1.5em;
          color: rgba(70, 42, 79, 0.75);
          margin: 0.25em 0 1em 0;
        }
        p {
          font-family: "IBM Plex Sans", sans-serif;
          font-size: 0.875em;
          line-height: 1.125em;
          color: rgba(70, 42, 79, 0.5);
        }
      }
      .illustration {
        img {
          width: 60vw;
          max-width: 420px;
        }
      }
    }
    .collegeList {
      display: flex;
      flex-direction: column;
      width: 45%;
      margin: 0 auto;
      padding: 0;
      font-family: "IBM Plex Sans", sans-serif;
      color: #462a4f;
      font-weight: 500;
      font-size: 24px;
      .name {
        text-align: left;
      }
      .value {
        text-align: right;
      }
      .collegeObj {
        border-radius: 4px;
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        background: #ffffff;
        border: 1px solid #eeeeee;
        box-sizing: border-box;
        //box-shadow: 0px 6px 12px rgba(51, 51, 51, 0.05);
        .top {
          display: flex;
          flex-direction: column;
          h3 {
            font-family: "IBM Plex Sans", sans-serif;
            color: #462a4f;
            font-weight: 600;
          }
          p {
            font-family: "IBM Plex Mono", monospace;
            font-size: 13px;
            line-height: 18px;
            margin-top: 4px;
            font-style: normal;
            font-weight: normal;
          }
          .divider {
            height: 0;
            width: 100%;
            border: 0.5px solid #eee;
            margin: 1em 0;
          }
        }
        .collegeDetails {
          display: flex;
          .section {
            display: flex;
            flex-direction: column;
            align-items: flex-start;
            h4 {
              font-family: "IBM Plex Sans", sans-serif;
              font-style: normal;
              letter-spacing: 0.05em;
              text-transform: uppercase;
              color: #666666;
            }
            p {
              border-radius: 2px;
              font-family: "IBM Plex Mono", monospace;
              font-style: normal;
              font-weight: 600;
            }
          }
        }
      }
      .good {
        background: #f1f9f1;
        color: #6aa565;
      }
      .bad {
        background: #ffedf1;
        color: #d53973;
      }
      .noupdate {
        background: #eee;
        color: #666;
      }
      .neutral {
        background: #fcf8ee;
        color: #d2a226;
      }
      .closed {
        color: #fd436d;
      }
      .opened {
        color: #6d72c5;
      }
      // loader
      .spinner {
        grid-column: 1 / span2;
        margin: 100px auto;
        width: 40px;
        height: 40px;
        position: relative;
        text-align: center;

        -webkit-animation: sk-rotate 2s infinite linear;
        animation: sk-rotate 2s infinite linear;
      }

      .dot1,
      .dot2 {
        width: 60%;
        height: 60%;
        display: inline-block;
        position: absolute;
        top: 0;
        background-color: #6d72c5;
        border-radius: 100%;

        -webkit-animation: sk-bounce 2s infinite ease-in-out;
        animation: sk-bounce 2s infinite ease-in-out;
      }

      .dot2 {
        top: auto;
        bottom: 0;
        -webkit-animation-delay: -1s;
        animation-delay: -1s;
      }

      @-webkit-keyframes sk-rotate {
        100% {
          -webkit-transform: rotate(360deg);
        }
      }
      @keyframes sk-rotate {
        100% {
          transform: rotate(360deg);
          -webkit-transform: rotate(360deg);
        }
      }

      @-webkit-keyframes sk-bounce {
        0%,
        100% {
          -webkit-transform: scale(0);
        }
        50% {
          -webkit-transform: scale(1);
        }
      }

      @keyframes sk-bounce {
        0%,
        100% {
          transform: scale(0);
          -webkit-transform: scale(0);
        }
        50% {
          transform: scale(1);
          -webkit-transform: scale(1);
        }
      }
    }
  }
}

#inputBox {
  font-size: 16px;
  padding-bottom: 10px;
}

#main {
  width:50%;
  display: flex;
  flex-direction: column;
  transform: translateY(120px);
}

.title {
  font-size: 32px;
  text-align: center;
  font-family: "IBM Plex Sans", sans-serif;
  width: 100%;
  transform: translateY(150px);
  color: black;
  font-weight: 400;
}

#tamjham {
  transform: translateY(200px);
}

.text {
  padding-bottom: 12px;
}

.name {
  text-align: left;
}

.value {
  text-align: right;
}
</style>
