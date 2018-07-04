import os, sys
sys.path.append(os.getcwd())
import aggregatedlog
logger = aggregatedlog.aggregatedlog.AggregatedLog()
print(logger.fetch_all())