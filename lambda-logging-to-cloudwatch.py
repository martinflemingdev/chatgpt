import logging
import json

# Configure logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)

def lambda_handler(event, context):
    try:
        # Log that the function is starting
        logger.info("Starting the Lambda function.")

        # Simulate some work here
        # For demonstration, we'll just echo back the received event
        result = {
            "statusCode": 200,
            "body": json.dumps({
                "message": "Success",
                "input": event
            })
        }

        # Log success message
        logger.info("The Lambda function has completed successfully.")

        return result
    except Exception as e:
        # Log any exceptions that occur
        logger.exception("An error occurred during the Lambda execution.")
        raise e  # Re-raise the exception to let Lambda handle it (e.g., for retries or DLQ)

# Example event for testing in AWS Lambda console or through an AWS SDK
if __name__ == "__main__":
    event = {"key": "value"}  # Sample event
    context = {}  # Placeholder for the context parameter
    print(lambda_handler(event, context))
