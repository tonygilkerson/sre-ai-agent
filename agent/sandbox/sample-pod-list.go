package main

const samplePodList2 = "This is a test"
const samplePodList = `
{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "annotations": {
                    "cni.projectcalico.org/containerID": "09f61657fe8579c433263f8a83d98be46cb583ccf0f2f26fd9364017acd179d5",
                    "cni.projectcalico.org/podIP": "10.1.7.18/32",
                    "cni.projectcalico.org/podIPs": "10.1.7.18/32"
                },
                "creationTimestamp": "2024-11-04T01:34:05Z",
                "generateName": "calico-kube-controllers-759cd8b574-",
                "labels": {
                    "k8s-app": "calico-kube-controllers",
                    "pod-template-hash": "759cd8b574"
                },
                "name": "calico-kube-controllers-759cd8b574-4wmb2",
                "namespace": "kube-system",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "calico-kube-controllers-759cd8b574",
                        "uid": "28e22f2f-a44b-4367-bb70-b82033e6225d"
                    }
                ],
                "resourceVersion": "51332441",
                "uid": "68187a2c-38a0-4e00-9862-916ff2036d05"
            },
            "spec": {
                "containers": [
                    {
                        "env": [
                            {
                                "name": "ENABLED_CONTROLLERS",
                                "value": "node"
                            },
                            {
                                "name": "DATASTORE_TYPE",
                                "value": "kubernetes"
                            }
                        ],
                        "image": "docker.io/calico/kube-controllers:v3.25.1",
                        "imagePullPolicy": "IfNotPresent",
                        "livenessProbe": {
                            "exec": {
                                "command": [
                                    "/usr/bin/check-status",
                                    "-l"
                                ]
                            },
                            "failureThreshold": 6,
                            "initialDelaySeconds": 10,
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 10
                        },
                        "name": "calico-kube-controllers",
                        "readinessProbe": {
                            "exec": {
                                "command": [
                                    "/usr/bin/check-status",
                                    "-r"
                                ]
                            },
                            "failureThreshold": 3,
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 1
                        },
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-q598t",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "nodeName": "zoo",
                "nodeSelector": {
                    "kubernetes.io/os": "linux"
                },
                "preemptionPolicy": "PreemptLowerPriority",
                "priority": 2000000000,
                "priorityClassName": "system-cluster-critical",
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "calico-kube-controllers",
                "serviceAccountName": "calico-kube-controllers",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "key": "CriticalAddonsOnly",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoSchedule",
                        "key": "node-role.kubernetes.io/master"
                    },
                    {
                        "effect": "NoSchedule",
                        "key": "node-role.kubernetes.io/control-plane"
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "name": "kube-api-access-q598t",
                        "projected": {
                            "defaultMode": 420,
                            "sources": [
                                {
                                    "serviceAccountToken": {
                                        "expirationSeconds": 3607,
                                        "path": "token"
                                    }
                                },
                                {
                                    "configMap": {
                                        "items": [
                                            {
                                                "key": "ca.crt",
                                                "path": "ca.crt"
                                            }
                                        ],
                                        "name": "kube-root-ca.crt"
                                    }
                                },
                                {
                                    "downwardAPI": {
                                        "items": [
                                            {
                                                "fieldRef": {
                                                    "apiVersion": "v1",
                                                    "fieldPath": "metadata.namespace"
                                                },
                                                "path": "namespace"
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:32:32Z",
                        "status": "True",
                        "type": "PodReadyToStartContainers"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2024-11-04T01:34:34Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:33:14Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:33:14Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2024-11-04T01:34:33Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "containerd://d4aef139fdc746c21bda0e9b51770b1faf9593fede211e547b458c1def289243",
                        "image": "docker.io/calico/kube-controllers:v3.25.1",
                        "imageID": "docker.io/calico/kube-controllers@sha256:02c1232ee4b8c5a145c401ac1adb34a63ee7fc46b70b6ad0a4e068a774f25f8a",
                        "lastState": {
                            "terminated": {
                                "containerID": "containerd://6975e218b508424246ecb3c9270cbfdf3705e6c201a2fe0498c905c723ad9939",
                                "exitCode": 255,
                                "finishedAt": "2025-06-10T19:31:25Z",
                                "reason": "Unknown",
                                "startedAt": "2025-05-21T01:49:13Z"
                            }
                        },
                        "name": "calico-kube-controllers",
                        "ready": true,
                        "restartCount": 17,
                        "started": true,
                        "state": {
                            "running": {
                                "startedAt": "2025-06-10T19:32:32Z"
                            }
                        },
                        "volumeMounts": [
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-q598t",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            }
                        ]
                    }
                ],
                "hostIP": "192.168.50.10",
                "hostIPs": [
                    {
                        "ip": "192.168.50.10"
                    }
                ],
                "phase": "Running",
                "podIP": "10.1.7.18",
                "podIPs": [
                    {
                        "ip": "10.1.7.18"
                    }
                ],
                "qosClass": "BestEffort",
                "startTime": "2024-11-04T01:34:34Z"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "creationTimestamp": "2025-04-22T14:51:17Z",
                "generateName": "calico-node-",
                "labels": {
                    "controller-revision-hash": "59c6b87f78",
                    "k8s-app": "calico-node",
                    "pod-template-generation": "1"
                },
                "name": "calico-node-f5rbq",
                "namespace": "kube-system",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "DaemonSet",
                        "name": "calico-node",
                        "uid": "eed6ab02-55bd-4260-a81f-5f91c7ab4421"
                    }
                ],
                "resourceVersion": "51332626",
                "uid": "2d14903c-9398-4462-9927-a34d499ba889"
            },
            "spec": {
                "affinity": {
                    "nodeAffinity": {
                        "requiredDuringSchedulingIgnoredDuringExecution": {
                            "nodeSelectorTerms": [
                                {
                                    "matchFields": [
                                        {
                                            "key": "metadata.name",
                                            "operator": "In",
                                            "values": [
                                                "zoo"
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    }
                },
                "containers": [
                    {
                        "env": [
                            {
                                "name": "DATASTORE_TYPE",
                                "value": "kubernetes"
                            },
                            {
                                "name": "WAIT_FOR_DATASTORE",
                                "value": "true"
                            },
                            {
                                "name": "NODENAME",
                                "valueFrom": {
                                    "fieldRef": {
                                        "apiVersion": "v1",
                                        "fieldPath": "spec.nodeName"
                                    }
                                }
                            },
                            {
                                "name": "CALICO_NETWORKING_BACKEND",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "calico_backend",
                                        "name": "calico-config"
                                    }
                                }
                            },
                            {
                                "name": "CLUSTER_TYPE",
                                "value": "k8s,bgp"
                            },
                            {
                                "name": "IP",
                                "value": "autodetect"
                            },
                            {
                                "name": "IP_AUTODETECTION_METHOD",
                                "value": "first-found"
                            },
                            {
                                "name": "CALICO_IPV4POOL_VXLAN",
                                "value": "Always"
                            },
                            {
                                "name": "CALICO_IPV6POOL_VXLAN",
                                "value": "Never"
                            },
                            {
                                "name": "FELIX_IPINIPMTU",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "veth_mtu",
                                        "name": "calico-config"
                                    }
                                }
                            },
                            {
                                "name": "FELIX_VXLANMTU",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "veth_mtu",
                                        "name": "calico-config"
                                    }
                                }
                            },
                            {
                                "name": "FELIX_WIREGUARDMTU",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "veth_mtu",
                                        "name": "calico-config"
                                    }
                                }
                            },
                            {
                                "name": "CALICO_IPV4POOL_CIDR",
                                "value": "10.1.0.0/16"
                            },
                            {
                                "name": "CALICO_DISABLE_FILE_LOGGING",
                                "value": "true"
                            },
                            {
                                "name": "FELIX_DEFAULTENDPOINTTOHOSTACTION",
                                "value": "ACCEPT"
                            },
                            {
                                "name": "FELIX_IPV6SUPPORT",
                                "value": "false"
                            },
                            {
                                "name": "FELIX_HEALTHENABLED",
                                "value": "true"
                            },
                            {
                                "name": "FELIX_FEATUREDETECTOVERRIDE",
                                "value": "ChecksumOffloadBroken=true"
                            }
                        ],
                        "envFrom": [
                            {
                                "configMapRef": {
                                    "name": "kubernetes-services-endpoint",
                                    "optional": true
                                }
                            }
                        ],
                        "image": "docker.io/calico/node:v3.25.1",
                        "imagePullPolicy": "IfNotPresent",
                        "lifecycle": {
                            "preStop": {
                                "exec": {
                                    "command": [
                                        "/bin/calico-node",
                                        "-shutdown"
                                    ]
                                }
                            }
                        },
                        "livenessProbe": {
                            "exec": {
                                "command": [
                                    "/bin/calico-node",
                                    "-felix-live"
                                ]
                            },
                            "failureThreshold": 6,
                            "initialDelaySeconds": 10,
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 10
                        },
                        "name": "calico-node",
                        "readinessProbe": {
                            "exec": {
                                "command": [
                                    "/bin/calico-node",
                                    "-felix-ready"
                                ]
                            },
                            "failureThreshold": 3,
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 10
                        },
                        "resources": {
                            "requests": {
                                "cpu": "250m"
                            }
                        },
                        "securityContext": {
                            "privileged": true
                        },
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/host/etc/cni/net.d",
                                "name": "cni-net-dir"
                            },
                            {
                                "mountPath": "/lib/modules",
                                "name": "lib-modules",
                                "readOnly": true
                            },
                            {
                                "mountPath": "/run/xtables.lock",
                                "name": "xtables-lock"
                            },
                            {
                                "mountPath": "/var/run/calico",
                                "name": "var-run-calico"
                            },
                            {
                                "mountPath": "/var/lib/calico",
                                "name": "var-lib-calico"
                            },
                            {
                                "mountPath": "/var/run/nodeagent",
                                "name": "policysync"
                            },
                            {
                                "mountPath": "/var/log/calico/cni",
                                "name": "cni-log-dir",
                                "readOnly": true
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-dd8rk",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "hostNetwork": true,
                "initContainers": [
                    {
                        "command": [
                            "/opt/cni/bin/calico-ipam",
                            "-upgrade"
                        ],
                        "env": [
                            {
                                "name": "KUBERNETES_NODE_NAME",
                                "valueFrom": {
                                    "fieldRef": {
                                        "apiVersion": "v1",
                                        "fieldPath": "spec.nodeName"
                                    }
                                }
                            },
                            {
                                "name": "CALICO_NETWORKING_BACKEND",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "calico_backend",
                                        "name": "calico-config"
                                    }
                                }
                            }
                        ],
                        "envFrom": [
                            {
                                "configMapRef": {
                                    "name": "kubernetes-services-endpoint",
                                    "optional": true
                                }
                            }
                        ],
                        "image": "docker.io/calico/cni:v3.25.1",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "upgrade-ipam",
                        "resources": {},
                        "securityContext": {
                            "privileged": true
                        },
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/var/lib/cni/networks",
                                "name": "host-local-net-dir"
                            },
                            {
                                "mountPath": "/host/opt/cni/bin",
                                "name": "cni-bin-dir"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-dd8rk",
                                "readOnly": true
                            }
                        ]
                    },
                    {
                        "command": [
                            "/opt/cni/bin/install"
                        ],
                        "env": [
                            {
                                "name": "CNI_CONF_NAME",
                                "value": "10-calico.conflist"
                            },
                            {
                                "name": "CNI_NETWORK_CONFIG",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "cni_network_config",
                                        "name": "calico-config"
                                    }
                                }
                            },
                            {
                                "name": "KUBERNETES_NODE_NAME",
                                "valueFrom": {
                                    "fieldRef": {
                                        "apiVersion": "v1",
                                        "fieldPath": "spec.nodeName"
                                    }
                                }
                            },
                            {
                                "name": "CNI_MTU",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "veth_mtu",
                                        "name": "calico-config"
                                    }
                                }
                            },
                            {
                                "name": "SLEEP",
                                "value": "false"
                            },
                            {
                                "name": "CNI_NET_DIR",
                                "value": "/var/snap/microk8s/current/args/cni-network"
                            }
                        ],
                        "envFrom": [
                            {
                                "configMapRef": {
                                    "name": "kubernetes-services-endpoint",
                                    "optional": true
                                }
                            }
                        ],
                        "image": "docker.io/calico/cni:v3.25.1",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "install-cni",
                        "resources": {},
                        "securityContext": {
                            "privileged": true
                        },
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/host/opt/cni/bin",
                                "name": "cni-bin-dir"
                            },
                            {
                                "mountPath": "/host/etc/cni/net.d",
                                "name": "cni-net-dir"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-dd8rk",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "nodeName": "zoo",
                "nodeSelector": {
                    "kubernetes.io/os": "linux"
                },
                "preemptionPolicy": "PreemptLowerPriority",
                "priority": 2000001000,
                "priorityClassName": "system-node-critical",
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "calico-node",
                "serviceAccountName": "calico-node",
                "terminationGracePeriodSeconds": 0,
                "tolerations": [
                    {
                        "effect": "NoSchedule",
                        "operator": "Exists"
                    },
                    {
                        "key": "CriticalAddonsOnly",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoExecute",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoSchedule",
                        "key": "node.kubernetes.io/disk-pressure",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoSchedule",
                        "key": "node.kubernetes.io/memory-pressure",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoSchedule",
                        "key": "node.kubernetes.io/pid-pressure",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoSchedule",
                        "key": "node.kubernetes.io/unschedulable",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoSchedule",
                        "key": "node.kubernetes.io/network-unavailable",
                        "operator": "Exists"
                    }
                ],
                "volumes": [
                    {
                        "hostPath": {
                            "path": "/lib/modules",
                            "type": ""
                        },
                        "name": "lib-modules"
                    },
                    {
                        "hostPath": {
                            "path": "/var/snap/microk8s/current/var/run/calico",
                            "type": ""
                        },
                        "name": "var-run-calico"
                    },
                    {
                        "hostPath": {
                            "path": "/var/snap/microk8s/current/var/lib/calico",
                            "type": ""
                        },
                        "name": "var-lib-calico"
                    },
                    {
                        "hostPath": {
                            "path": "/run/xtables.lock",
                            "type": "FileOrCreate"
                        },
                        "name": "xtables-lock"
                    },
                    {
                        "hostPath": {
                            "path": "/sys/fs/",
                            "type": "DirectoryOrCreate"
                        },
                        "name": "sys-fs"
                    },
                    {
                        "hostPath": {
                            "path": "/var/snap/microk8s/current/opt/cni/bin",
                            "type": ""
                        },
                        "name": "cni-bin-dir"
                    },
                    {
                        "hostPath": {
                            "path": "/var/snap/microk8s/current/args/cni-network",
                            "type": ""
                        },
                        "name": "cni-net-dir"
                    },
                    {
                        "hostPath": {
                            "path": "/var/snap/microk8s/common/var/log/calico/cni",
                            "type": ""
                        },
                        "name": "cni-log-dir"
                    },
                    {
                        "hostPath": {
                            "path": "/var/snap/microk8s/current/var/lib/cni/networks",
                            "type": ""
                        },
                        "name": "host-local-net-dir"
                    },
                    {
                        "hostPath": {
                            "path": "/var/snap/microk8s/current/var/run/nodeagent",
                            "type": "DirectoryOrCreate"
                        },
                        "name": "policysync"
                    },
                    {
                        "name": "kube-api-access-dd8rk",
                        "projected": {
                            "defaultMode": 420,
                            "sources": [
                                {
                                    "serviceAccountToken": {
                                        "expirationSeconds": 3607,
                                        "path": "token"
                                    }
                                },
                                {
                                    "configMap": {
                                        "items": [
                                            {
                                                "key": "ca.crt",
                                                "path": "ca.crt"
                                            }
                                        ],
                                        "name": "kube-root-ca.crt"
                                    }
                                },
                                {
                                    "downwardAPI": {
                                        "items": [
                                            {
                                                "fieldRef": {
                                                    "apiVersion": "v1",
                                                    "fieldPath": "metadata.namespace"
                                                },
                                                "path": "namespace"
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:31:56Z",
                        "status": "True",
                        "type": "PodReadyToStartContainers"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-04-22T14:51:27Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:33:55Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:33:55Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-04-22T14:51:19Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "containerd://524623a5c3f2553816d7397b6768b12edf5811be4bb10ecc650287e2fac7e2cc",
                        "image": "docker.io/calico/node:v3.25.1",
                        "imageID": "docker.io/calico/node@sha256:0cd00e83d06b3af8cd712ad2c310be07b240235ad7ca1397e04eb14d20dcc20f",
                        "lastState": {
                            "terminated": {
                                "containerID": "containerd://5e94222e19d9e44b94a4fc5ed3cfb533ac50cf47e6d586b958875a85a9a0aa2f",
                                "exitCode": 255,
                                "finishedAt": "2025-06-10T19:31:24Z",
                                "reason": "Unknown",
                                "startedAt": "2025-05-21T01:50:16Z"
                            }
                        },
                        "name": "calico-node",
                        "ready": true,
                        "restartCount": 3,
                        "started": true,
                        "state": {
                            "running": {
                                "startedAt": "2025-06-10T19:33:32Z"
                            }
                        },
                        "volumeMounts": [
                            {
                                "mountPath": "/host/etc/cni/net.d",
                                "name": "cni-net-dir"
                            },
                            {
                                "mountPath": "/lib/modules",
                                "name": "lib-modules",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            },
                            {
                                "mountPath": "/run/xtables.lock",
                                "name": "xtables-lock"
                            },
                            {
                                "mountPath": "/var/run/calico",
                                "name": "var-run-calico"
                            },
                            {
                                "mountPath": "/var/lib/calico",
                                "name": "var-lib-calico"
                            },
                            {
                                "mountPath": "/var/run/nodeagent",
                                "name": "policysync"
                            },
                            {
                                "mountPath": "/var/log/calico/cni",
                                "name": "cni-log-dir",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-dd8rk",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            }
                        ]
                    }
                ],
                "hostIP": "192.168.50.10",
                "hostIPs": [
                    {
                        "ip": "192.168.50.10"
                    }
                ],
                "initContainerStatuses": [
                    {
                        "containerID": "containerd://425bd8ae2898244b6a5bc07842706dbce376b4b133beb088ef870b0a9719a2b7",
                        "image": "docker.io/calico/cni:v3.25.1",
                        "imageID": "docker.io/calico/cni@sha256:9a2c99f0314053aa11e971bd5d72e17951767bf5c6ff1fd9c38c4582d7cb8a0a",
                        "lastState": {},
                        "name": "upgrade-ipam",
                        "ready": true,
                        "restartCount": 3,
                        "started": false,
                        "state": {
                            "terminated": {
                                "containerID": "containerd://425bd8ae2898244b6a5bc07842706dbce376b4b133beb088ef870b0a9719a2b7",
                                "exitCode": 0,
                                "finishedAt": "2025-06-10T19:32:00Z",
                                "reason": "Completed",
                                "startedAt": "2025-06-10T19:31:56Z"
                            }
                        },
                        "volumeMounts": [
                            {
                                "mountPath": "/var/lib/cni/networks",
                                "name": "host-local-net-dir"
                            },
                            {
                                "mountPath": "/host/opt/cni/bin",
                                "name": "cni-bin-dir"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-dd8rk",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            }
                        ]
                    },
                    {
                        "containerID": "containerd://d1cd084f226a9323f57fd9fda2da4d793d0f470f821c853ef062061f93ccb109",
                        "image": "docker.io/calico/cni:v3.25.1",
                        "imageID": "docker.io/calico/cni@sha256:9a2c99f0314053aa11e971bd5d72e17951767bf5c6ff1fd9c38c4582d7cb8a0a",
                        "lastState": {},
                        "name": "install-cni",
                        "ready": true,
                        "restartCount": 0,
                        "started": false,
                        "state": {
                            "terminated": {
                                "containerID": "containerd://d1cd084f226a9323f57fd9fda2da4d793d0f470f821c853ef062061f93ccb109",
                                "exitCode": 0,
                                "finishedAt": "2025-06-10T19:33:27Z",
                                "reason": "Completed",
                                "startedAt": "2025-06-10T19:32:03Z"
                            }
                        },
                        "volumeMounts": [
                            {
                                "mountPath": "/host/opt/cni/bin",
                                "name": "cni-bin-dir"
                            },
                            {
                                "mountPath": "/host/etc/cni/net.d",
                                "name": "cni-net-dir"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-dd8rk",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            }
                        ]
                    }
                ],
                "phase": "Running",
                "podIP": "192.168.50.10",
                "podIPs": [
                    {
                        "ip": "192.168.50.10"
                    }
                ],
                "qosClass": "Burstable",
                "startTime": "2025-04-22T14:51:19Z"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "annotations": {
                    "cni.projectcalico.org/containerID": "15754e56b17a26ff91a9efec485831fa430b070a6fdccd0de0a2b394fc0cba4a",
                    "cni.projectcalico.org/podIP": "10.1.7.28/32",
                    "cni.projectcalico.org/podIPs": "10.1.7.28/32",
                    "priorityClassName": "system-cluster-critical"
                },
                "creationTimestamp": "2024-11-04T01:34:05Z",
                "generateName": "coredns-7896dbf49-",
                "labels": {
                    "k8s-app": "kube-dns",
                    "pod-template-hash": "7896dbf49"
                },
                "name": "coredns-7896dbf49-tmx5q",
                "namespace": "kube-system",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "coredns-7896dbf49",
                        "uid": "788f2af0-8807-44b6-bca5-8a61ed1a60fb"
                    }
                ],
                "resourceVersion": "51332428",
                "uid": "da8634a6-6954-4e77-9ff0-c2de03d2c120"
            },
            "spec": {
                "containers": [
                    {
                        "args": [
                            "-conf",
                            "/etc/coredns/Corefile"
                        ],
                        "image": "coredns/coredns:1.10.1",
                        "imagePullPolicy": "IfNotPresent",
                        "livenessProbe": {
                            "failureThreshold": 5,
                            "httpGet": {
                                "path": "/health",
                                "port": 8080,
                                "scheme": "HTTP"
                            },
                            "initialDelaySeconds": 60,
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 5
                        },
                        "name": "coredns",
                        "ports": [
                            {
                                "containerPort": 53,
                                "name": "dns",
                                "protocol": "UDP"
                            },
                            {
                                "containerPort": 53,
                                "name": "dns-tcp",
                                "protocol": "TCP"
                            },
                            {
                                "containerPort": 9153,
                                "name": "metrics",
                                "protocol": "TCP"
                            }
                        ],
                        "readinessProbe": {
                            "failureThreshold": 3,
                            "httpGet": {
                                "path": "/ready",
                                "port": 8181,
                                "scheme": "HTTP"
                            },
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 1
                        },
                        "resources": {
                            "limits": {
                                "memory": "170Mi"
                            },
                            "requests": {
                                "cpu": "100m",
                                "memory": "70Mi"
                            }
                        },
                        "securityContext": {
                            "allowPrivilegeEscalation": false,
                            "capabilities": {
                                "add": [
                                    "NET_BIND_SERVICE"
                                ],
                                "drop": [
                                    "all"
                                ]
                            },
                            "readOnlyRootFilesystem": true
                        },
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/etc/coredns",
                                "name": "config-volume",
                                "readOnly": true
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-qgcxf",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "Default",
                "enableServiceLinks": true,
                "nodeName": "zoo",
                "preemptionPolicy": "PreemptLowerPriority",
                "priority": 2000000000,
                "priorityClassName": "system-cluster-critical",
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "coredns",
                "serviceAccountName": "coredns",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "key": "CriticalAddonsOnly",
                        "operator": "Exists"
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "configMap": {
                            "defaultMode": 420,
                            "items": [
                                {
                                    "key": "Corefile",
                                    "path": "Corefile"
                                }
                            ],
                            "name": "coredns"
                        },
                        "name": "config-volume"
                    },
                    {
                        "name": "kube-api-access-qgcxf",
                        "projected": {
                            "defaultMode": 420,
                            "sources": [
                                {
                                    "serviceAccountToken": {
                                        "expirationSeconds": 3607,
                                        "path": "token"
                                    }
                                },
                                {
                                    "configMap": {
                                        "items": [
                                            {
                                                "key": "ca.crt",
                                                "path": "ca.crt"
                                            }
                                        ],
                                        "name": "kube-root-ca.crt"
                                    }
                                },
                                {
                                    "downwardAPI": {
                                        "items": [
                                            {
                                                "fieldRef": {
                                                    "apiVersion": "v1",
                                                    "fieldPath": "metadata.namespace"
                                                },
                                                "path": "namespace"
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:32:30Z",
                        "status": "True",
                        "type": "PodReadyToStartContainers"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2024-11-04T01:34:34Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:33:11Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:33:11Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2024-11-04T01:34:33Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "containerd://d21d508fa537366649067f7005d45501046ebd0fe47a8fd18800e57287deab90",
                        "image": "docker.io/coredns/coredns:1.10.1",
                        "imageID": "docker.io/coredns/coredns@sha256:a0ead06651cf580044aeb0a0feba63591858fb2e43ade8c9dea45a6a89ae7e5e",
                        "lastState": {
                            "terminated": {
                                "containerID": "containerd://7cf9a035988e59e995eaf7560b056fbd9fc74e0eaec56e24e691340510171654",
                                "exitCode": 255,
                                "finishedAt": "2025-06-10T19:31:25Z",
                                "reason": "Unknown",
                                "startedAt": "2025-05-21T01:49:18Z"
                            }
                        },
                        "name": "coredns",
                        "ready": true,
                        "restartCount": 17,
                        "started": true,
                        "state": {
                            "running": {
                                "startedAt": "2025-06-10T19:32:29Z"
                            }
                        },
                        "volumeMounts": [
                            {
                                "mountPath": "/etc/coredns",
                                "name": "config-volume",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-qgcxf",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            }
                        ]
                    }
                ],
                "hostIP": "192.168.50.10",
                "hostIPs": [
                    {
                        "ip": "192.168.50.10"
                    }
                ],
                "phase": "Running",
                "podIP": "10.1.7.28",
                "podIPs": [
                    {
                        "ip": "10.1.7.28"
                    }
                ],
                "qosClass": "Burstable",
                "startTime": "2024-11-04T01:34:34Z"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "annotations": {
                    "cni.projectcalico.org/containerID": "0ad21ce067c231ef5f3ed5c4b3e3cde41119461c8037f2746044fdb7b3955245",
                    "cni.projectcalico.org/podIP": "10.1.7.9/32",
                    "cni.projectcalico.org/podIPs": "10.1.7.9/32"
                },
                "creationTimestamp": "2024-11-04T01:46:24Z",
                "generateName": "hostpath-provisioner-5fbc49d86c-",
                "labels": {
                    "k8s-app": "hostpath-provisioner",
                    "pod-template-hash": "5fbc49d86c"
                },
                "name": "hostpath-provisioner-5fbc49d86c-9cqt2",
                "namespace": "kube-system",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "hostpath-provisioner-5fbc49d86c",
                        "uid": "4092dcb6-7f58-487f-b224-7eda61f47576"
                    }
                ],
                "resourceVersion": "51332304",
                "uid": "2e20f049-d000-49c6-8b14-5db0921b7ab2"
            },
            "spec": {
                "containers": [
                    {
                        "env": [
                            {
                                "name": "NAMESPACE",
                                "valueFrom": {
                                    "fieldRef": {
                                        "apiVersion": "v1",
                                        "fieldPath": "metadata.namespace"
                                    }
                                }
                            },
                            {
                                "name": "NODE_NAME",
                                "valueFrom": {
                                    "fieldRef": {
                                        "apiVersion": "v1",
                                        "fieldPath": "spec.nodeName"
                                    }
                                }
                            },
                            {
                                "name": "PV_DIR",
                                "value": "/var/snap/microk8s/common/default-storage"
                            },
                            {
                                "name": "BUSYBOX_IMAGE",
                                "value": "busybox:1.28.4"
                            }
                        ],
                        "image": "cdkbot/hostpath-provisioner:1.5.0",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "hostpath-provisioner",
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/var/snap/microk8s/common/default-storage",
                                "name": "pv-volume"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-zgfvf",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "nodeName": "zoo",
                "preemptionPolicy": "PreemptLowerPriority",
                "priority": 0,
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "microk8s-hostpath",
                "serviceAccountName": "microk8s-hostpath",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "hostPath": {
                            "path": "/var/snap/microk8s/common/default-storage",
                            "type": ""
                        },
                        "name": "pv-volume"
                    },
                    {
                        "name": "kube-api-access-zgfvf",
                        "projected": {
                            "defaultMode": 420,
                            "sources": [
                                {
                                    "serviceAccountToken": {
                                        "expirationSeconds": 3607,
                                        "path": "token"
                                    }
                                },
                                {
                                    "configMap": {
                                        "items": [
                                            {
                                                "key": "ca.crt",
                                                "path": "ca.crt"
                                            }
                                        ],
                                        "name": "kube-root-ca.crt"
                                    }
                                },
                                {
                                    "downwardAPI": {
                                        "items": [
                                            {
                                                "fieldRef": {
                                                    "apiVersion": "v1",
                                                    "fieldPath": "metadata.namespace"
                                                },
                                                "path": "namespace"
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:32:25Z",
                        "status": "True",
                        "type": "PodReadyToStartContainers"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2024-11-04T01:46:24Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:32:25Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:32:25Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2024-11-04T01:46:24Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "containerd://d41dc19b2f94a7fd8be2cf13ca4055e76c7df8c862d4f5a9954ef9d31fcf2a7e",
                        "image": "docker.io/cdkbot/hostpath-provisioner:1.5.0",
                        "imageID": "docker.io/cdkbot/hostpath-provisioner@sha256:ac51e50e32b70e47077fe90928a7fe4d3fc8dd49192db4932c2643c49729c2eb",
                        "lastState": {
                            "terminated": {
                                "containerID": "containerd://ef40c0a33bdeafef39373d12b14083f56bae212cb8f77ca7acb692fc270f5fe0",
                                "exitCode": 255,
                                "finishedAt": "2025-06-10T19:31:28Z",
                                "reason": "Unknown",
                                "startedAt": "2025-05-21T01:49:15Z"
                            }
                        },
                        "name": "hostpath-provisioner",
                        "ready": true,
                        "restartCount": 49,
                        "started": true,
                        "state": {
                            "running": {
                                "startedAt": "2025-06-10T19:32:24Z"
                            }
                        },
                        "volumeMounts": [
                            {
                                "mountPath": "/var/snap/microk8s/common/default-storage",
                                "name": "pv-volume"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-zgfvf",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            }
                        ]
                    }
                ],
                "hostIP": "192.168.50.10",
                "hostIPs": [
                    {
                        "ip": "192.168.50.10"
                    }
                ],
                "phase": "Running",
                "podIP": "10.1.7.9",
                "podIPs": [
                    {
                        "ip": "10.1.7.9"
                    }
                ],
                "qosClass": "BestEffort",
                "startTime": "2024-11-04T01:46:24Z"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "annotations": {
                    "cni.projectcalico.org/containerID": "8054610ee75654cb726f26ddf1d2de2b83815c2bee26e8ac3d318b4492d498cc",
                    "cni.projectcalico.org/podIP": "10.1.7.44/32",
                    "cni.projectcalico.org/podIPs": "10.1.7.44/32"
                },
                "creationTimestamp": "2024-11-04T01:46:26Z",
                "generateName": "metrics-server-df8dbf7f5-",
                "labels": {
                    "k8s-app": "metrics-server",
                    "pod-template-hash": "df8dbf7f5"
                },
                "name": "metrics-server-df8dbf7f5-dhf8q",
                "namespace": "kube-system",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "metrics-server-df8dbf7f5",
                        "uid": "998c309c-ba77-469a-81de-d7836b80bed2"
                    }
                ],
                "resourceVersion": "51332469",
                "uid": "68ad26e6-2069-4123-9637-9f1ec7ae7bc6"
            },
            "spec": {
                "containers": [
                    {
                        "args": [
                            "--cert-dir=/tmp",
                            "--secure-port=4443",
                            "--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname",
                            "--kubelet-use-node-status-port",
                            "--metric-resolution=15s",
                            "--kubelet-insecure-tls"
                        ],
                        "image": "registry.k8s.io/metrics-server/metrics-server:v0.6.3",
                        "imagePullPolicy": "IfNotPresent",
                        "livenessProbe": {
                            "failureThreshold": 3,
                            "httpGet": {
                                "path": "/livez",
                                "port": "https",
                                "scheme": "HTTPS"
                            },
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 1
                        },
                        "name": "metrics-server",
                        "ports": [
                            {
                                "containerPort": 4443,
                                "name": "https",
                                "protocol": "TCP"
                            }
                        ],
                        "readinessProbe": {
                            "failureThreshold": 3,
                            "httpGet": {
                                "path": "/readyz",
                                "port": "https",
                                "scheme": "HTTPS"
                            },
                            "initialDelaySeconds": 20,
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 1
                        },
                        "resources": {
                            "requests": {
                                "cpu": "100m",
                                "memory": "200Mi"
                            }
                        },
                        "securityContext": {
                            "readOnlyRootFilesystem": true,
                            "runAsNonRoot": true,
                            "runAsUser": 1000
                        },
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/tmp",
                                "name": "tmp-dir"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-zn48x",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "nodeName": "zoo",
                "nodeSelector": {
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/os": "linux"
                },
                "preemptionPolicy": "PreemptLowerPriority",
                "priority": 2000000000,
                "priorityClassName": "system-cluster-critical",
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "metrics-server",
                "serviceAccountName": "metrics-server",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "emptyDir": {},
                        "name": "tmp-dir"
                    },
                    {
                        "name": "kube-api-access-zn48x",
                        "projected": {
                            "defaultMode": 420,
                            "sources": [
                                {
                                    "serviceAccountToken": {
                                        "expirationSeconds": 3607,
                                        "path": "token"
                                    }
                                },
                                {
                                    "configMap": {
                                        "items": [
                                            {
                                                "key": "ca.crt",
                                                "path": "ca.crt"
                                            }
                                        ],
                                        "name": "kube-root-ca.crt"
                                    }
                                },
                                {
                                    "downwardAPI": {
                                        "items": [
                                            {
                                                "fieldRef": {
                                                    "apiVersion": "v1",
                                                    "fieldPath": "metadata.namespace"
                                                },
                                                "path": "namespace"
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:32:21Z",
                        "status": "True",
                        "type": "PodReadyToStartContainers"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2024-11-04T01:46:27Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:33:19Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2025-06-10T19:33:19Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2024-11-04T01:46:26Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "containerd://b78a06a29b2d8e06d5b1bd50dfa199db18f498b357ca1244114204a8b21a9622",
                        "image": "registry.k8s.io/metrics-server/metrics-server:v0.6.3",
                        "imageID": "registry.k8s.io/metrics-server/metrics-server@sha256:c60778fa1c44d0c5a0c4530ebe83f9243ee6fc02f4c3dc59226c201931350b10",
                        "lastState": {
                            "terminated": {
                                "containerID": "containerd://cae9780a93f9922bd9b828ac473d771570d34688416e2082ecd4b6536115b6b4",
                                "exitCode": 2,
                                "finishedAt": "2025-06-10T19:32:47Z",
                                "reason": "Error",
                                "startedAt": "2025-06-10T19:32:20Z"
                            }
                        },
                        "name": "metrics-server",
                        "ready": true,
                        "restartCount": 29,
                        "started": true,
                        "state": {
                            "running": {
                                "startedAt": "2025-06-10T19:32:58Z"
                            }
                        },
                        "volumeMounts": [
                            {
                                "mountPath": "/tmp",
                                "name": "tmp-dir"
                            },
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-zn48x",
                                "readOnly": true,
                                "recursiveReadOnly": "Disabled"
                            }
                        ]
                    }
                ],
                "hostIP": "192.168.50.10",
                "hostIPs": [
                    {
                        "ip": "192.168.50.10"
                    }
                ],
                "phase": "Running",
                "podIP": "10.1.7.44",
                "podIPs": [
                    {
                        "ip": "10.1.7.44"
                    }
                ],
                "qosClass": "Burstable",
                "startTime": "2024-11-04T01:46:27Z"
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": ""
    }
}
`
