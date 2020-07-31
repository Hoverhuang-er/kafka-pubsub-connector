kafka-pubsub-connector
============

Description
-----------

Kafka to Google PubSub or AWS SNS or AWS SQS or Azure ServiceBus connector.

Install
-------

Directly from Go:

.. code:: bash

    $ make release

Or locally, in development mode:

.. code:: bash

    $ git clone https://github.com/Hoverhuang-er/kafka-pubsub-connector.git
    $ cd kafka-pubsub-connector
    $ make prepare

Or Docker mode:

.. code:: bash

    docker build -t kafka-pubsub-connector:latest .
    docker run -rm -e BROKER="" \
               -e ZOOKEEPER=""  \
               -e TOPIC=""      \
               -e PROVIDER=""   \
               -e CRED=""  kafka-pubsub-connector:latest

Reference
---------

- `<https://gocloud.dev/>`_
- `<https://cloud.google.com/pubsub/docs/reference/libraries#client-libraries-install-python>`_