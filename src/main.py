from neural_network.train_neural_network import train
from neural_network.classify import classify
from alfred import Alfred
import os.path
import logging
import sys

def main():
    if not os.path.isfile("synapses.json"):
        logging.info("Training the bot...")
        train(hidden_neurons=20, alpha=0.1, epochs=100000, dropout=False, dropout_percent=0.2)
    if len(sys.argv) < 2:
        alfred = Alfred("cli")
	while True:
            data = raw_input("Ask me anything: ")
            print(classify(data))
    else:
        alfred = Alfred(sys.argv[1])
    return

main()
