#!/bin/bash

echo "🔧 Step-by-step Slack App Setup"

# Step 1: Start Slack App Creation
echo ""
echo "📌 Step 1: Visit the link below to create a new app:"
echo "   https://api.slack.com/apps?new_app=1"
read -p "Press Enter after you create the app..."

# Step 2: Gather App Info
read -p "Enter your new Slack App Client ID: " CLIENT_ID
read -p "Enter your new Slack App Client Secret: " CLIENT_SECRET

# Step 3: Add Redirect URI
echo ""
echo "➡️  Add this redirect URL to your app settings under 'OAuth & Permissions':"
echo "   http://localhost:3000/oauth/callback"
read -p "Press Enter after adding the redirect URI..."

# Step 4: Add Scopes
echo ""
echo "➕ Add this User Token Scope:"
echo "   users.profile:write"
read -p "Press Enter after adding scopes and reinstalling the app..."

# Step 5: Save to .env or shell
echo ""
echo "✅ Slack app is ready!"
echo ""
echo "You can now add this to your environment or .env file:"
echo ""
echo "export SLACK_CLIENT_ID=$CLIENT_ID"
echo "export SLACK_CLIENT_SECRET=$CLIENT_SECRET"
