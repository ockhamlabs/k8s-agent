package labels

const (
	HeliosSpot            = "scheduling.helios/spot"
	HeliosSpotFallback    = "scheduling.helios/spot-fallback"
	HeliosFakeSpot        = "scheduling.helios/fake-spot"
	KopsSpot              = "spot"
	KarpenterCapacityType = "karpenter.sh/capacity-type"
	WorkerSpot            = "node-role.kubernetes.io/spot-worker"

	ValueKarpenterCapacityTypeSpot     = "spot"
	ValueKarpenterCapacityTypeOnDemand = "on-demand"
)
