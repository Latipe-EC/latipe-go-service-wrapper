@echo off
echo run server file

rem Array of service paths
set services=latipe-email-service/ latipe-delivery-service/ latipe-promotion-service/ order-service-v2/ latipe-transaction-service/  latipe-notification-service/

rem Create new tabs and run commands
for %%s in (%services%) do (
   start cmd /c "cd %%s && make startw"
)


