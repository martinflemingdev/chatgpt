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


#########################################################################################
    
import logging
import json

# Configure logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)

def lambda_handler(event, context):
    try:
        # Log that the function is starting
        logger.info("Starting the Lambda function.")

        # Here, you might perform your Lambda function's main work
        # This is just a placeholder for demonstration
        # For example, you might attempt to open a non-existent file or divide by zero

        # Simulate an error
        1 / 0  # This will cause a ZeroDivisionError

        # If the function reaches this point, log a success message
        logger.info("The Lambda function has completed successfully.")
        
        # Return a successful response
        return {
            "statusCode": 200,
            "body": json.dumps({"message": "Success"})
        }
    except Exception as e:
        # Log the exception with its stack trace
        logger.exception("An error occurred during the Lambda execution.")

        # Instead of re-raising the exception, return an error response or perform other cleanup
        return {
            "statusCode": 500,
            "body": json.dumps({"message": "An error occurred"})
        }

# Example event for testing in AWS Lambda console or through an AWS SDK
if __name__ == "__main__":
    event = {"key": "value"}  # Sample event
    context = {}  # Placeholder for the context parameter
    print(lambda_handler(event, context))
