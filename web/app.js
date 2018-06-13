new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        ReceivedContent: '',// Holds new messages to be sent to the server
        Content: '', // Holds new messages to be sent to the server
        UserName: null,
        joined: false, // True if email and username have been filled in
        GroupName: '',
        Status: null,
        Action: '',
        OnlineUsers: '',
        GroupMember: [],
    },

    created: function () {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function (e) {
            var msg = JSON.parse(e.data);
                if (msg.Action == "GetUsers") {
                    for (var i = 0; i < msg.GroupMember.length; i++) {
                        self.OnlineUsers += '<div class="white-text">' + msg.GroupMember[i] + '</div>' + '<br/>';
                    }
                } else if (msg.Action == "SendMessageTo") {
                    self.ReceivedContent +=
                        '<div class="chip">' + msg.UserName + ' from: ' + msg.GroupName + '</div>' +
                        '<div class="white-text">' + msg.Content + '</div>' +
                        '<br/>';
                }
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight;// Auto scroll to the bottom
        });
    },

    methods: {
        send: function () {
            if (this.Content != '') {
                this.ws.send(
                    JSON.stringify({
                        UserName: this.UserName,
                        Content: $('<p>').html(this.Content).text(), // Strip out html
                        GroupName: this.GroupName,
                        Action: "SendMessageTo"
                    }));
                this.Content = ''; // Reset newMsg
            }
        },

        join: function () {
            if (!this.UserName) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.UserName = $('<p>').html(this.UserName).text();
            this.ws.send(
              JSON.stringify({
                  UserName: this.UserName,
                  Action: "GetUsers"
              })
            );
           /* this.ws.send(
                JSON.stringify({
                    UserName: this.UserName,
                    Action: "CreateUser"
                }));*/
            this.joined = true;
        },
    }
});