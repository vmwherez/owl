// Cody, write a server in C++ please:

#include <iostream>
#include <string>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>

using namespace std;

int main()
{
    int sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd < 0)
    {
        cout << "Error opening socket" << endl;
        return 1;
    }

    struct sockaddr_in server_addr;
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(8080);
    inet_pton(AF_INET, "127.0.0.1", &server_addr.sin_addr);

    if (bind(sockfd, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0)
    {
        cout << "Error binding socket" << endl;
        return 1;
    }

    if (listen(sockfd, 10) < 0)
    {
        cout << "Error listening on socket" << endl;
        return 1;
    }

    cout << "Listening on port 8080" << endl;

    while (true)
    {
        struct sockaddr_in client_addr;
        socklen_t client_addr_len = sizeof(client_addr);
        int client_sockfd = accept(sockfd, (struct sockaddr *)&client_addr, &client_addr_len);
        if (client_sockfd < 0)
        {
            cout << "Error accepting connection" << endl;
            return 1;
        }

        char buffer[1024];
        int bytes_read = read(client_sockfd, buffer, 1024);
        if (bytes_read < 0)
        {
            cout << "Error reading from socket" << endl;
            return 1;
        }

        cout << "Received: " << buffer << endl;

        string response = "HTTP/1.1 200 OK\nContent-Type: text/html\n\n<html><body><h1>Hello, world!</h1></body></html>";

        if (write(client_sockfd, response.c_str(), response.size()) < 0)
        {
            cout << "Error writing to socket" << endl;
            return 1;
        }

        close(client_sockfd);
        cout << "Closed connection" << endl;

        cout << "Waiting for next connection" << endl;

        usleep(1000000);

        cout << "Done" << endl;
    }
