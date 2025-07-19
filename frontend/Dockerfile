# Use an official Node.js runtime as a parent image.
# Using alpine for a smaller image size.
FROM node:20-alpine

# Update packages to patch vulnerabilities
RUN apk update && apk upgrade

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to leverage Docker cache
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application's source code
COPY . .

# The command to run the development server.
# The `--host` flag is crucial to expose the server outside the container.
CMD ["npm", "run", "dev", "--", "--host"]