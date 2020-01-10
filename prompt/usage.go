package prompt

const PROMPT_USAGE =
`
   cf prompt [--logged-out "logged out format string"] "logged in format string"

PLACEHOLDERS:
   %a	The currently targeted API endpoint
   %o	The currently targeted org
   %s	The currently targeted space
   %u	The currently logged in user
   %v	The version of the API
`
