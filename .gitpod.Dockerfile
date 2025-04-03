FROM gitpod/workspace-full

USER gitpod

# Install Go, MongoDB server, and MongoDB client
RUN sudo apt-get update && sudo apt-get install -y \
    golang \
    mongodb \
    libmongo-client-dev

# Start MongoDB service in the background
RUN sudo service mongod start

# Expose the ports for MongoDB and application
EXPOSE 27017
EXPOSE 8080

# Set the default command for the container
CMD ["bash"]
