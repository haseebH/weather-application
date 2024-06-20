# Use an official Node.js runtime as a parent image
FROM node:22-alpine

# Set the working directory in the container
WORKDIR /app

# Copy package.json and yarn.lock
COPY package.json yarn.lock ./

# Install dependencies using yarn
RUN yarn install

# Copy the rest of the application code
COPY . .

# Build the React app
RUN yarn build

# Serve the React app using a static file server
# Install `serve` globally
RUN yarn global add serve

# Expose the port the app runs on
EXPOSE 8082

# Command to run the application
CMD ["yarn", "start"]
