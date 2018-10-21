import neural_network as nn
import json
import numpy as np
import sys

def classify(sentence, show_details=False):
    ERROR_THRESHOLD = 0.2
    # load our calculated synapse values
    synapse_file = 'synapses.json' 
    with open(synapse_file) as data_file: 
        synapse = json.load(data_file) 
        synapse_0 = np.asarray(synapse['synapse0']) 
        synapse_1 = np.asarray(synapse['synapse1'])
        words = synapse['words']
        classes = synapse['classes']

    results = nn.think(sentence, synapse_0, synapse_1, words, show_details)

    results = [[i,r] for i,r in enumerate(results) if r>ERROR_THRESHOLD ] 
    results.sort(key=lambda x: x[1], reverse=True) 
    return_results =[[classes[r[0]],r[1]] for r in results]
    print ("%s" % (return_results[0][0]))
    return return_results
