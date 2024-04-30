package resources

// AAQCRDs is a map containing yaml strings of all CRDs
var AAQCRDs map[string]string = map[string]string{
	"aaq": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: aaqs.aaq.kubevirt.io
spec:
  group: aaq.kubevirt.io
  names:
    kind: AAQ
    listKind: AAQList
    plural: aaqs
    shortNames:
    - aaq
    - aaqs
    singular: aaq
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    - jsonPath: .status.phase
      name: Phase
      type: string
    name: v1alpha1
    schema:
      openAPIV3Schema:
        description: AAQ is the AAQ Operator CRD
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AAQSpec defines our specification for the AAQ installation
            properties:
              certConfig:
                description: certificate configuration
                properties:
                  ca:
                    description: CA configuration CA certs are kept in the CA bundle
                      as long as they are valid
                    properties:
                      duration:
                        description: The requested 'duration' (i.e. lifetime) of the
                          Certificate.
                        type: string
                      renewBefore:
                        description: The amount of time before the currently issued
                          certificate's ` + "`" + `notAfter` + "`" + ` time that we will begin to attempt
                          to renew the certificate.
                        type: string
                    type: object
                  server:
                    description: Server configuration Certs are rotated and discarded
                    properties:
                      duration:
                        description: The requested 'duration' (i.e. lifetime) of the
                          Certificate.
                        type: string
                      renewBefore:
                        description: The amount of time before the currently issued
                          certificate's ` + "`" + `notAfter` + "`" + ` time that we will begin to attempt
                          to renew the certificate.
                        type: string
                    type: object
                type: object
              configuration:
                description: holds aaq configurations.
                properties:
                  allowApplicationAwareClusterResourceQuota:
                    description: AllowApplicationAwareClusterResourceQuota can be
                      set to true to allow creation and management of ApplicationAwareClusterResourceQuota.
                      Defaults to false
                    type: boolean
                  sidecarEvaluators:
                    description: SidecarEvaluators allow custom quota counting for
                      external operator
                    items:
                      description: A single application container that you want to
                        run within a pod.
                      properties:
                        args:
                          description: 'Arguments to the entrypoint. The container
                            image''s CMD is used if this is not provided. Variable
                            references $(VAR_NAME) are expanded using the container''s
                            environment. If a variable cannot be resolved, the reference
                            in the input string will be unchanged. Double $$ are reduced
                            to a single $, which allows for escaping the $(VAR_NAME)
                            syntax: i.e. "$$(VAR_NAME)" will produce the string literal
                            "$(VAR_NAME)". Escaped references will never be expanded,
                            regardless of whether the variable exists or not. Cannot
                            be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell'
                          items:
                            type: string
                          type: array
                        command:
                          description: 'Entrypoint array. Not executed within a shell.
                            The container image''s ENTRYPOINT is used if this is not
                            provided. Variable references $(VAR_NAME) are expanded
                            using the container''s environment. If a variable cannot
                            be resolved, the reference in the input string will be
                            unchanged. Double $$ are reduced to a single $, which
                            allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)"
                            will produce the string literal "$(VAR_NAME)". Escaped
                            references will never be expanded, regardless of whether
                            the variable exists or not. Cannot be updated. More info:
                            https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell'
                          items:
                            type: string
                          type: array
                        env:
                          description: List of environment variables to set in the
                            container. Cannot be updated.
                          items:
                            description: EnvVar represents an environment variable
                              present in a Container.
                            properties:
                              name:
                                description: Name of the environment variable. Must
                                  be a C_IDENTIFIER.
                                type: string
                              value:
                                description: 'Variable references $(VAR_NAME) are
                                  expanded using the previously defined environment
                                  variables in the container and any service environment
                                  variables. If a variable cannot be resolved, the
                                  reference in the input string will be unchanged.
                                  Double $$ are reduced to a single $, which allows
                                  for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)"
                                  will produce the string literal "$(VAR_NAME)". Escaped
                                  references will never be expanded, regardless of
                                  whether the variable exists or not. Defaults to
                                  "".'
                                type: string
                              valueFrom:
                                description: Source for the environment variable's
                                  value. Cannot be used if value is not empty.
                                properties:
                                  configMapKeyRef:
                                    description: Selects a key of a ConfigMap.
                                    properties:
                                      key:
                                        description: The key to select.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info:
                                          https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                          TODO: Add other useful fields. apiVersion,
                                          kind, uid?'
                                        type: string
                                      optional:
                                        description: Specify whether the ConfigMap
                                          or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                    x-kubernetes-map-type: atomic
                                  fieldRef:
                                    description: 'Selects a field of the pod: supports
                                      metadata.name, metadata.namespace, ` + "`" + `metadata.labels[''<KEY>'']` + "`" + `,
                                      ` + "`" + `metadata.annotations[''<KEY>'']` + "`" + `, spec.nodeName,
                                      spec.serviceAccountName, status.hostIP, status.podIP,
                                      status.podIPs.'
                                    properties:
                                      apiVersion:
                                        description: Version of the schema the FieldPath
                                          is written in terms of, defaults to "v1".
                                        type: string
                                      fieldPath:
                                        description: Path of the field to select in
                                          the specified API version.
                                        type: string
                                    required:
                                    - fieldPath
                                    type: object
                                    x-kubernetes-map-type: atomic
                                  resourceFieldRef:
                                    description: 'Selects a resource of the container:
                                      only resources limits and requests (limits.cpu,
                                      limits.memory, limits.ephemeral-storage, requests.cpu,
                                      requests.memory and requests.ephemeral-storage)
                                      are currently supported.'
                                    properties:
                                      containerName:
                                        description: 'Container name: required for
                                          volumes, optional for env vars'
                                        type: string
                                      divisor:
                                        anyOf:
                                        - type: integer
                                        - type: string
                                        description: Specifies the output format of
                                          the exposed resources, defaults to "1"
                                        pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                        x-kubernetes-int-or-string: true
                                      resource:
                                        description: 'Required: resource to select'
                                        type: string
                                    required:
                                    - resource
                                    type: object
                                    x-kubernetes-map-type: atomic
                                  secretKeyRef:
                                    description: Selects a key of a secret in the
                                      pod's namespace
                                    properties:
                                      key:
                                        description: The key of the secret to select
                                          from.  Must be a valid secret key.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info:
                                          https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                          TODO: Add other useful fields. apiVersion,
                                          kind, uid?'
                                        type: string
                                      optional:
                                        description: Specify whether the Secret or
                                          its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                    x-kubernetes-map-type: atomic
                                type: object
                            required:
                            - name
                            type: object
                          type: array
                        envFrom:
                          description: List of sources to populate environment variables
                            in the container. The keys defined within a source must
                            be a C_IDENTIFIER. All invalid keys will be reported as
                            an event when the container is starting. When a key exists
                            in multiple sources, the value associated with the last
                            source will take precedence. Values defined by an Env
                            with a duplicate key will take precedence. Cannot be updated.
                          items:
                            description: EnvFromSource represents the source of a
                              set of ConfigMaps
                            properties:
                              configMapRef:
                                description: The ConfigMap to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info:
                                      https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                      TODO: Add other useful fields. apiVersion, kind,
                                      uid?'
                                    type: string
                                  optional:
                                    description: Specify whether the ConfigMap must
                                      be defined
                                    type: boolean
                                type: object
                                x-kubernetes-map-type: atomic
                              prefix:
                                description: An optional identifier to prepend to
                                  each key in the ConfigMap. Must be a C_IDENTIFIER.
                                type: string
                              secretRef:
                                description: The Secret to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info:
                                      https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                      TODO: Add other useful fields. apiVersion, kind,
                                      uid?'
                                    type: string
                                  optional:
                                    description: Specify whether the Secret must be
                                      defined
                                    type: boolean
                                type: object
                                x-kubernetes-map-type: atomic
                            type: object
                          type: array
                        image:
                          description: 'Container image name. More info: https://kubernetes.io/docs/concepts/containers/images
                            This field is optional to allow higher level config management
                            to default or override container images in workload controllers
                            like Deployments and StatefulSets.'
                          type: string
                        imagePullPolicy:
                          description: 'Image pull policy. One of Always, Never, IfNotPresent.
                            Defaults to Always if :latest tag is specified, or IfNotPresent
                            otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images'
                          type: string
                        lifecycle:
                          description: Actions that the management system should take
                            in response to container lifecycle events. Cannot be updated.
                          properties:
                            postStart:
                              description: 'PostStart is called immediately after
                                a container is created. If the handler fails, the
                                container is terminated and restarted according to
                                its restart policy. Other management of the container
                                blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks'
                              properties:
                                exec:
                                  description: Exec specifies the action to take.
                                  properties:
                                    command:
                                      description: Command is the command line to
                                        execute inside the container, the working
                                        directory for the command  is root ('/') in
                                        the container's filesystem. The command is
                                        simply exec'd, it is not run inside a shell,
                                        so traditional shell instructions ('|', etc)
                                        won't work. To use a shell, you need to explicitly
                                        call out to that shell. Exit status of 0 is
                                        treated as live/healthy and non-zero is unhealthy.
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request
                                    to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults
                                        to the pod IP. You probably want to set "Host"
                                        in httpHeaders instead.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request.
                                        HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom
                                          header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name. This
                                              will be canonicalized upon output, so
                                              case-variant names will be understood
                                              as the same header.
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access
                                        on the container. Number must be in the range
                                        1 to 65535. Name must be an IANA_SVC_NAME.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to
                                        the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: Deprecated. TCPSocket is NOT supported
                                    as a LifecycleHandler and kept for the backward
                                    compatibility. There are no validation of this
                                    field and lifecycle hooks will fail in runtime
                                    when tcp handler is specified.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect
                                        to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access
                                        on the container. Number must be in the range
                                        1 to 65535. Name must be an IANA_SVC_NAME.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                            preStop:
                              description: 'PreStop is called immediately before a
                                container is terminated due to an API request or management
                                event such as liveness/startup probe failure, preemption,
                                resource contention, etc. The handler is not called
                                if the container crashes or exits. The Pod''s termination
                                grace period countdown begins before the PreStop hook
                                is executed. Regardless of the outcome of the handler,
                                the container will eventually terminate within the
                                Pod''s termination grace period (unless delayed by
                                finalizers). Other management of the container blocks
                                until the hook completes or until the termination
                                grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks'
                              properties:
                                exec:
                                  description: Exec specifies the action to take.
                                  properties:
                                    command:
                                      description: Command is the command line to
                                        execute inside the container, the working
                                        directory for the command  is root ('/') in
                                        the container's filesystem. The command is
                                        simply exec'd, it is not run inside a shell,
                                        so traditional shell instructions ('|', etc)
                                        won't work. To use a shell, you need to explicitly
                                        call out to that shell. Exit status of 0 is
                                        treated as live/healthy and non-zero is unhealthy.
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request
                                    to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults
                                        to the pod IP. You probably want to set "Host"
                                        in httpHeaders instead.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request.
                                        HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom
                                          header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name. This
                                              will be canonicalized upon output, so
                                              case-variant names will be understood
                                              as the same header.
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access
                                        on the container. Number must be in the range
                                        1 to 65535. Name must be an IANA_SVC_NAME.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to
                                        the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: Deprecated. TCPSocket is NOT supported
                                    as a LifecycleHandler and kept for the backward
                                    compatibility. There are no validation of this
                                    field and lifecycle hooks will fail in runtime
                                    when tcp handler is specified.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect
                                        to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access
                                        on the container. Number must be in the range
                                        1 to 65535. Name must be an IANA_SVC_NAME.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                          type: object
                        livenessProbe:
                          description: 'Periodic probe of container liveness. Container
                            will be restarted if the probe fails. Cannot be updated.
                            More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                          properties:
                            exec:
                              description: Exec specifies the action to take.
                              properties:
                                command:
                                  description: Command is the command line to execute
                                    inside the container, the working directory for
                                    the command  is root ('/') in the container's
                                    filesystem. The command is simply exec'd, it is
                                    not run inside a shell, so traditional shell instructions
                                    ('|', etc) won't work. To use a shell, you need
                                    to explicitly call out to that shell. Exit status
                                    of 0 is treated as live/healthy and non-zero is
                                    unhealthy.
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: Minimum consecutive failures for the probe
                                to be considered failed after having succeeded. Defaults
                                to 3. Minimum value is 1.
                              format: int32
                              type: integer
                            grpc:
                              description: GRPC specifies an action involving a GRPC
                                port.
                              properties:
                                port:
                                  description: Port number of the gRPC service. Number
                                    must be in the range 1 to 65535.
                                  format: int32
                                  type: integer
                                service:
                                  description: "Service is the name of the service
                                    to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
                                    \n If this is not specified, the default behavior
                                    is defined by gRPC."
                                  type: string
                              required:
                              - port
                              type: object
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to
                                    the pod IP. You probably want to set "Host" in
                                    httpHeaders instead.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request.
                                    HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header
                                      to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name. This will
                                          be canonicalized upon output, so case-variant
                                          names will be understood as the same header.
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access
                                    on the container. Number must be in the range
                                    1 to 65535. Name must be an IANA_SVC_NAME.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the
                                    host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: 'Number of seconds after the container
                                has started before liveness probes are initiated.
                                More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe.
                                Default to 10 seconds. Minimum value is 1.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe
                                to be considered successful after having failed. Defaults
                                to 1. Must be 1 for liveness and startup. Minimum
                                value is 1.
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving
                                a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to,
                                    defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access
                                    on the container. Number must be in the range
                                    1 to 65535. Name must be an IANA_SVC_NAME.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs
                                to terminate gracefully upon probe failure. The grace
                                period is the duration in seconds after the processes
                                running in the pod are sent a termination signal and
                                the time when the processes are forcibly halted with
                                a kill signal. Set this value longer than the expected
                                cleanup time for your process. If this value is nil,
                                the pod's terminationGracePeriodSeconds will be used.
                                Otherwise, this value overrides the value provided
                                by the pod spec. Value must be non-negative integer.
                                The value zero indicates stop immediately via the
                                kill signal (no opportunity to shut down). This is
                                a beta field and requires enabling ProbeTerminationGracePeriod
                                feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds
                                is used if unset.
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: 'Number of seconds after which the probe
                                times out. Defaults to 1 second. Minimum value is
                                1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                              format: int32
                              type: integer
                          type: object
                        name:
                          description: Name of the container specified as a DNS_LABEL.
                            Each container in a pod must have a unique name (DNS_LABEL).
                            Cannot be updated.
                          type: string
                        ports:
                          description: List of ports to expose from the container.
                            Not specifying a port here DOES NOT prevent that port
                            from being exposed. Any port which is listening on the
                            default "0.0.0.0" address inside a container will be accessible
                            from the network. Modifying this array with strategic
                            merge patch may corrupt the data. For more information
                            See https://github.com/kubernetes/kubernetes/issues/108255.
                            Cannot be updated.
                          items:
                            description: ContainerPort represents a network port in
                              a single container.
                            properties:
                              containerPort:
                                description: Number of port to expose on the pod's
                                  IP address. This must be a valid port number, 0
                                  < x < 65536.
                                format: int32
                                type: integer
                              hostIP:
                                description: What host IP to bind the external port
                                  to.
                                type: string
                              hostPort:
                                description: Number of port to expose on the host.
                                  If specified, this must be a valid port number,
                                  0 < x < 65536. If HostNetwork is specified, this
                                  must match ContainerPort. Most containers do not
                                  need this.
                                format: int32
                                type: integer
                              name:
                                description: If specified, this must be an IANA_SVC_NAME
                                  and unique within the pod. Each named port in a
                                  pod must have a unique name. Name for the port that
                                  can be referred to by services.
                                type: string
                              protocol:
                                default: TCP
                                description: Protocol for port. Must be UDP, TCP,
                                  or SCTP. Defaults to "TCP".
                                type: string
                            required:
                            - containerPort
                            type: object
                          type: array
                          x-kubernetes-list-map-keys:
                          - containerPort
                          - protocol
                          x-kubernetes-list-type: map
                        readinessProbe:
                          description: 'Periodic probe of container service readiness.
                            Container will be removed from service endpoints if the
                            probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                          properties:
                            exec:
                              description: Exec specifies the action to take.
                              properties:
                                command:
                                  description: Command is the command line to execute
                                    inside the container, the working directory for
                                    the command  is root ('/') in the container's
                                    filesystem. The command is simply exec'd, it is
                                    not run inside a shell, so traditional shell instructions
                                    ('|', etc) won't work. To use a shell, you need
                                    to explicitly call out to that shell. Exit status
                                    of 0 is treated as live/healthy and non-zero is
                                    unhealthy.
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: Minimum consecutive failures for the probe
                                to be considered failed after having succeeded. Defaults
                                to 3. Minimum value is 1.
                              format: int32
                              type: integer
                            grpc:
                              description: GRPC specifies an action involving a GRPC
                                port.
                              properties:
                                port:
                                  description: Port number of the gRPC service. Number
                                    must be in the range 1 to 65535.
                                  format: int32
                                  type: integer
                                service:
                                  description: "Service is the name of the service
                                    to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
                                    \n If this is not specified, the default behavior
                                    is defined by gRPC."
                                  type: string
                              required:
                              - port
                              type: object
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to
                                    the pod IP. You probably want to set "Host" in
                                    httpHeaders instead.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request.
                                    HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header
                                      to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name. This will
                                          be canonicalized upon output, so case-variant
                                          names will be understood as the same header.
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access
                                    on the container. Number must be in the range
                                    1 to 65535. Name must be an IANA_SVC_NAME.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the
                                    host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: 'Number of seconds after the container
                                has started before liveness probes are initiated.
                                More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe.
                                Default to 10 seconds. Minimum value is 1.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe
                                to be considered successful after having failed. Defaults
                                to 1. Must be 1 for liveness and startup. Minimum
                                value is 1.
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving
                                a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to,
                                    defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access
                                    on the container. Number must be in the range
                                    1 to 65535. Name must be an IANA_SVC_NAME.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs
                                to terminate gracefully upon probe failure. The grace
                                period is the duration in seconds after the processes
                                running in the pod are sent a termination signal and
                                the time when the processes are forcibly halted with
                                a kill signal. Set this value longer than the expected
                                cleanup time for your process. If this value is nil,
                                the pod's terminationGracePeriodSeconds will be used.
                                Otherwise, this value overrides the value provided
                                by the pod spec. Value must be non-negative integer.
                                The value zero indicates stop immediately via the
                                kill signal (no opportunity to shut down). This is
                                a beta field and requires enabling ProbeTerminationGracePeriod
                                feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds
                                is used if unset.
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: 'Number of seconds after which the probe
                                times out. Defaults to 1 second. Minimum value is
                                1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                              format: int32
                              type: integer
                          type: object
                        resizePolicy:
                          description: Resources resize policy for the container.
                          items:
                            description: ContainerResizePolicy represents resource
                              resize policy for the container.
                            properties:
                              resourceName:
                                description: 'Name of the resource to which this resource
                                  resize policy applies. Supported values: cpu, memory.'
                                type: string
                              restartPolicy:
                                description: Restart policy to apply when specified
                                  resource is resized. If not specified, it defaults
                                  to NotRequired.
                                type: string
                            required:
                            - resourceName
                            - restartPolicy
                            type: object
                          type: array
                          x-kubernetes-list-type: atomic
                        resources:
                          description: 'Compute Resources required by this container.
                            Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                          properties:
                            claims:
                              description: "Claims lists the names of resources, defined
                                in spec.resourceClaims, that are used by this container.
                                \n This is an alpha field and requires enabling the
                                DynamicResourceAllocation feature gate. \n This field
                                is immutable. It can only be set for containers."
                              items:
                                description: ResourceClaim references one entry in
                                  PodSpec.ResourceClaims.
                                properties:
                                  name:
                                    description: Name must match the name of one entry
                                      in pod.spec.resourceClaims of the Pod where
                                      this field is used. It makes that resource available
                                      inside a container.
                                    type: string
                                required:
                                - name
                                type: object
                              type: array
                              x-kubernetes-list-map-keys:
                              - name
                              x-kubernetes-list-type: map
                            limits:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: 'Limits describes the maximum amount of
                                compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                              type: object
                            requests:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: 'Requests describes the minimum amount
                                of compute resources required. If Requests is omitted
                                for a container, it defaults to Limits if that is
                                explicitly specified, otherwise to an implementation-defined
                                value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                              type: object
                          type: object
                        restartPolicy:
                          description: 'RestartPolicy defines the restart behavior
                            of individual containers in a pod. This field may only
                            be set for init containers, and the only allowed value
                            is "Always". For non-init containers or when this field
                            is not specified, the restart behavior is defined by the
                            Pod''s restart policy and the container type. Setting
                            the RestartPolicy as "Always" for the init container will
                            have the following effect: this init container will be
                            continually restarted on exit until all regular containers
                            have terminated. Once all regular containers have completed,
                            all init containers with restartPolicy "Always" will be
                            shut down. This lifecycle differs from normal init containers
                            and is often referred to as a "sidecar" container. Although
                            this init container still starts in the init container
                            sequence, it does not wait for the container to complete
                            before proceeding to the next init container. Instead,
                            the next init container starts immediately after this
                            init container is started, or after any startupProbe has
                            successfully completed.'
                          type: string
                        securityContext:
                          description: 'SecurityContext defines the security options
                            the container should be run with. If set, the fields of
                            SecurityContext override the equivalent fields of PodSecurityContext.
                            More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/'
                          properties:
                            allowPrivilegeEscalation:
                              description: 'AllowPrivilegeEscalation controls whether
                                a process can gain more privileges than its parent
                                process. This bool directly controls if the no_new_privs
                                flag will be set on the container process. AllowPrivilegeEscalation
                                is true always when the container is: 1) run as Privileged
                                2) has CAP_SYS_ADMIN Note that this field cannot be
                                set when spec.os.name is windows.'
                              type: boolean
                            capabilities:
                              description: The capabilities to add/drop when running
                                containers. Defaults to the default set of capabilities
                                granted by the container runtime. Note that this field
                                cannot be set when spec.os.name is windows.
                              properties:
                                add:
                                  description: Added capabilities
                                  items:
                                    description: Capability represent POSIX capabilities
                                      type
                                    type: string
                                  type: array
                                drop:
                                  description: Removed capabilities
                                  items:
                                    description: Capability represent POSIX capabilities
                                      type
                                    type: string
                                  type: array
                              type: object
                            privileged:
                              description: Run container in privileged mode. Processes
                                in privileged containers are essentially equivalent
                                to root on the host. Defaults to false. Note that
                                this field cannot be set when spec.os.name is windows.
                              type: boolean
                            procMount:
                              description: procMount denotes the type of proc mount
                                to use for the containers. The default is DefaultProcMount
                                which uses the container runtime defaults for readonly
                                paths and masked paths. This requires the ProcMountType
                                feature flag to be enabled. Note that this field cannot
                                be set when spec.os.name is windows.
                              type: string
                            readOnlyRootFilesystem:
                              description: Whether this container has a read-only
                                root filesystem. Default is false. Note that this
                                field cannot be set when spec.os.name is windows.
                              type: boolean
                            runAsGroup:
                              description: The GID to run the entrypoint of the container
                                process. Uses runtime default if unset. May also be
                                set in PodSecurityContext.  If set in both SecurityContext
                                and PodSecurityContext, the value specified in SecurityContext
                                takes precedence. Note that this field cannot be set
                                when spec.os.name is windows.
                              format: int64
                              type: integer
                            runAsNonRoot:
                              description: Indicates that the container must run as
                                a non-root user. If true, the Kubelet will validate
                                the image at runtime to ensure that it does not run
                                as UID 0 (root) and fail to start the container if
                                it does. If unset or false, no such validation will
                                be performed. May also be set in PodSecurityContext.  If
                                set in both SecurityContext and PodSecurityContext,
                                the value specified in SecurityContext takes precedence.
                              type: boolean
                            runAsUser:
                              description: The UID to run the entrypoint of the container
                                process. Defaults to user specified in image metadata
                                if unspecified. May also be set in PodSecurityContext.  If
                                set in both SecurityContext and PodSecurityContext,
                                the value specified in SecurityContext takes precedence.
                                Note that this field cannot be set when spec.os.name
                                is windows.
                              format: int64
                              type: integer
                            seLinuxOptions:
                              description: The SELinux context to be applied to the
                                container. If unspecified, the container runtime will
                                allocate a random SELinux context for each container.  May
                                also be set in PodSecurityContext.  If set in both
                                SecurityContext and PodSecurityContext, the value
                                specified in SecurityContext takes precedence. Note
                                that this field cannot be set when spec.os.name is
                                windows.
                              properties:
                                level:
                                  description: Level is SELinux level label that applies
                                    to the container.
                                  type: string
                                role:
                                  description: Role is a SELinux role label that applies
                                    to the container.
                                  type: string
                                type:
                                  description: Type is a SELinux type label that applies
                                    to the container.
                                  type: string
                                user:
                                  description: User is a SELinux user label that applies
                                    to the container.
                                  type: string
                              type: object
                            seccompProfile:
                              description: The seccomp options to use by this container.
                                If seccomp options are provided at both the pod &
                                container level, the container options override the
                                pod options. Note that this field cannot be set when
                                spec.os.name is windows.
                              properties:
                                localhostProfile:
                                  description: localhostProfile indicates a profile
                                    defined in a file on the node should be used.
                                    The profile must be preconfigured on the node
                                    to work. Must be a descending path, relative to
                                    the kubelet's configured seccomp profile location.
                                    Must be set if type is "Localhost". Must NOT be
                                    set for any other type.
                                  type: string
                                type:
                                  description: "type indicates which kind of seccomp
                                    profile will be applied. Valid options are: \n
                                    Localhost - a profile defined in a file on the
                                    node should be used. RuntimeDefault - the container
                                    runtime default profile should be used. Unconfined
                                    - no profile should be applied."
                                  type: string
                              required:
                              - type
                              type: object
                            windowsOptions:
                              description: The Windows specific settings applied to
                                all containers. If unspecified, the options from the
                                PodSecurityContext will be used. If set in both SecurityContext
                                and PodSecurityContext, the value specified in SecurityContext
                                takes precedence. Note that this field cannot be set
                                when spec.os.name is linux.
                              properties:
                                gmsaCredentialSpec:
                                  description: GMSACredentialSpec is where the GMSA
                                    admission webhook (https://github.com/kubernetes-sigs/windows-gmsa)
                                    inlines the contents of the GMSA credential spec
                                    named by the GMSACredentialSpecName field.
                                  type: string
                                gmsaCredentialSpecName:
                                  description: GMSACredentialSpecName is the name
                                    of the GMSA credential spec to use.
                                  type: string
                                hostProcess:
                                  description: HostProcess determines if a container
                                    should be run as a 'Host Process' container. All
                                    of a Pod's containers must have the same effective
                                    HostProcess value (it is not allowed to have a
                                    mix of HostProcess containers and non-HostProcess
                                    containers). In addition, if HostProcess is true
                                    then HostNetwork must also be set to true.
                                  type: boolean
                                runAsUserName:
                                  description: The UserName in Windows to run the
                                    entrypoint of the container process. Defaults
                                    to the user specified in image metadata if unspecified.
                                    May also be set in PodSecurityContext. If set
                                    in both SecurityContext and PodSecurityContext,
                                    the value specified in SecurityContext takes precedence.
                                  type: string
                              type: object
                          type: object
                        startupProbe:
                          description: 'StartupProbe indicates that the Pod has successfully
                            initialized. If specified, no other probes are executed
                            until this completes successfully. If this probe fails,
                            the Pod will be restarted, just as if the livenessProbe
                            failed. This can be used to provide different probe parameters
                            at the beginning of a Pod''s lifecycle, when it might
                            take a long time to load data or warm a cache, than during
                            steady-state operation. This cannot be updated. More info:
                            https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                          properties:
                            exec:
                              description: Exec specifies the action to take.
                              properties:
                                command:
                                  description: Command is the command line to execute
                                    inside the container, the working directory for
                                    the command  is root ('/') in the container's
                                    filesystem. The command is simply exec'd, it is
                                    not run inside a shell, so traditional shell instructions
                                    ('|', etc) won't work. To use a shell, you need
                                    to explicitly call out to that shell. Exit status
                                    of 0 is treated as live/healthy and non-zero is
                                    unhealthy.
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: Minimum consecutive failures for the probe
                                to be considered failed after having succeeded. Defaults
                                to 3. Minimum value is 1.
                              format: int32
                              type: integer
                            grpc:
                              description: GRPC specifies an action involving a GRPC
                                port.
                              properties:
                                port:
                                  description: Port number of the gRPC service. Number
                                    must be in the range 1 to 65535.
                                  format: int32
                                  type: integer
                                service:
                                  description: "Service is the name of the service
                                    to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
                                    \n If this is not specified, the default behavior
                                    is defined by gRPC."
                                  type: string
                              required:
                              - port
                              type: object
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to
                                    the pod IP. You probably want to set "Host" in
                                    httpHeaders instead.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request.
                                    HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header
                                      to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name. This will
                                          be canonicalized upon output, so case-variant
                                          names will be understood as the same header.
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access
                                    on the container. Number must be in the range
                                    1 to 65535. Name must be an IANA_SVC_NAME.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the
                                    host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: 'Number of seconds after the container
                                has started before liveness probes are initiated.
                                More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe.
                                Default to 10 seconds. Minimum value is 1.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe
                                to be considered successful after having failed. Defaults
                                to 1. Must be 1 for liveness and startup. Minimum
                                value is 1.
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving
                                a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to,
                                    defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access
                                    on the container. Number must be in the range
                                    1 to 65535. Name must be an IANA_SVC_NAME.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs
                                to terminate gracefully upon probe failure. The grace
                                period is the duration in seconds after the processes
                                running in the pod are sent a termination signal and
                                the time when the processes are forcibly halted with
                                a kill signal. Set this value longer than the expected
                                cleanup time for your process. If this value is nil,
                                the pod's terminationGracePeriodSeconds will be used.
                                Otherwise, this value overrides the value provided
                                by the pod spec. Value must be non-negative integer.
                                The value zero indicates stop immediately via the
                                kill signal (no opportunity to shut down). This is
                                a beta field and requires enabling ProbeTerminationGracePeriod
                                feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds
                                is used if unset.
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: 'Number of seconds after which the probe
                                times out. Defaults to 1 second. Minimum value is
                                1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes'
                              format: int32
                              type: integer
                          type: object
                        stdin:
                          description: Whether this container should allocate a buffer
                            for stdin in the container runtime. If this is not set,
                            reads from stdin in the container will always result in
                            EOF. Default is false.
                          type: boolean
                        stdinOnce:
                          description: Whether the container runtime should close
                            the stdin channel after it has been opened by a single
                            attach. When stdin is true the stdin stream will remain
                            open across multiple attach sessions. If stdinOnce is
                            set to true, stdin is opened on container start, is empty
                            until the first client attaches to stdin, and then remains
                            open and accepts data until the client disconnects, at
                            which time stdin is closed and remains closed until the
                            container is restarted. If this flag is false, a container
                            processes that reads from stdin will never receive an
                            EOF. Default is false
                          type: boolean
                        terminationMessagePath:
                          description: 'Optional: Path at which the file to which
                            the container''s termination message will be written is
                            mounted into the container''s filesystem. Message written
                            is intended to be brief final status, such as an assertion
                            failure message. Will be truncated by the node if greater
                            than 4096 bytes. The total message length across all containers
                            will be limited to 12kb. Defaults to /dev/termination-log.
                            Cannot be updated.'
                          type: string
                        terminationMessagePolicy:
                          description: Indicate how the termination message should
                            be populated. File will use the contents of terminationMessagePath
                            to populate the container status message on both success
                            and failure. FallbackToLogsOnError will use the last chunk
                            of container log output if the termination message file
                            is empty and the container exited with an error. The log
                            output is limited to 2048 bytes or 80 lines, whichever
                            is smaller. Defaults to File. Cannot be updated.
                          type: string
                        tty:
                          description: Whether this container should allocate a TTY
                            for itself, also requires 'stdin' to be true. Default
                            is false.
                          type: boolean
                        volumeDevices:
                          description: volumeDevices is the list of block devices
                            to be used by the container.
                          items:
                            description: volumeDevice describes a mapping of a raw
                              block device within a container.
                            properties:
                              devicePath:
                                description: devicePath is the path inside of the
                                  container that the device will be mapped to.
                                type: string
                              name:
                                description: name must match the name of a persistentVolumeClaim
                                  in the pod
                                type: string
                            required:
                            - devicePath
                            - name
                            type: object
                          type: array
                        volumeMounts:
                          description: Pod volumes to mount into the container's filesystem.
                            Cannot be updated.
                          items:
                            description: VolumeMount describes a mounting of a Volume
                              within a container.
                            properties:
                              mountPath:
                                description: Path within the container at which the
                                  volume should be mounted.  Must not contain ':'.
                                type: string
                              mountPropagation:
                                description: mountPropagation determines how mounts
                                  are propagated from the host to container and the
                                  other way around. When not set, MountPropagationNone
                                  is used. This field is beta in 1.10.
                                type: string
                              name:
                                description: This must match the Name of a Volume.
                                type: string
                              readOnly:
                                description: Mounted read-only if true, read-write
                                  otherwise (false or unspecified). Defaults to false.
                                type: boolean
                              subPath:
                                description: Path within the volume from which the
                                  container's volume should be mounted. Defaults to
                                  "" (volume's root).
                                type: string
                              subPathExpr:
                                description: Expanded path within the volume from
                                  which the container's volume should be mounted.
                                  Behaves similarly to SubPath but environment variable
                                  references $(VAR_NAME) are expanded using the container's
                                  environment. Defaults to "" (volume's root). SubPathExpr
                                  and SubPath are mutually exclusive.
                                type: string
                            required:
                            - mountPath
                            - name
                            type: object
                          type: array
                        workingDir:
                          description: Container's working directory. If not specified,
                            the container runtime's default will be used, which might
                            be configured in the container image. Cannot be updated.
                          type: string
                      required:
                      - name
                      type: object
                    type: array
                  vmiCalculatorConfiguration:
                    description: VmiCalculatorConfiguration determine how resource
                      allocation will be done with ApplicationAwareResourceQuota
                    properties:
                      configName:
                        default: DedicatedVirtualResources
                        description: 'ConfigName determine how resource allocation
                          will be done with ApplicationAwareResourceQuota. allowed
                          values are: VmiPodUsage, VirtualResources, DedicatedVirtualResources
                          or IgnoreVmiCalculator'
                        enum:
                        - VmiPodUsage
                        - VirtualResources
                        - DedicatedVirtualResources
                        - IgnoreVmiCalculator
                        type: string
                    type: object
                type: object
              imagePullPolicy:
                description: PullPolicy describes a policy for if/when to pull a container
                  image
                enum:
                - Always
                - IfNotPresent
                - Never
                type: string
              infra:
                description: Rules on which nodes AAQ infrastructure pods will be
                  scheduled
                properties:
                  affinity:
                    description: affinity enables pod affinity/anti-affinity placement
                      expanding the types of constraints that can be expressed with
                      nodeSelector. affinity is going to be applied to the relevant
                      kind of pods in parallel with nodeSelector See https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for
                          the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node matches the corresponding matchExpressions;
                              the node(s) with the highest sum are the most preferred.
                            items:
                              description: An empty preferred scheduling term matches
                                all objects with implicit weight 0 (i.e. it's a no-op).
                                A null preferred scheduling term matches no objects
                                (i.e. is also a no-op).
                              properties:
                                preference:
                                  description: A node selector term, associated with
                                    the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                weight:
                                  description: Weight associated with matching the
                                    corresponding nodeSelectorTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to an update), the system
                              may or may not try to eventually evict the pod from
                              its node.
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms.
                                  The terms are ORed.
                                items:
                                  description: A null or empty node selector term
                                    matches no objects. The requirements of them are
                                    ANDed. The TopologySelectorTerm type implements
                                    a subset of the NodeSelectorTerm.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                            x-kubernetes-map-type: atomic
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                          co-locate this pod in the same node, zone, etc. as some
                          other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to a pod label update),
                              the system may or may not try to eventually evict the
                              pod from its node. When there are multiple elements,
                              the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules
                          (e.g. avoid putting this pod in the same node, zone, etc.
                          as some other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the anti-affinity expressions
                              specified by this field, but it may choose a node that
                              violates one or more of the expressions. The node that
                              is most preferred is the one with the greatest sum of
                              weights, i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              anti-affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified
                              by this field are not met at scheduling time, the pod
                              will not be scheduled onto the node. If the anti-affinity
                              requirements specified by this field cease to be met
                              at some point during pod execution (e.g. due to a pod
                              label update), the system may or may not try to eventually
                              evict the pod from its node. When there are multiple
                              elements, the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: 'nodeSelector is the node selector applied to the
                      relevant kind of pods It specifies a map of key-value pairs:
                      for the pod to be eligible to run on a node, the node must have
                      each of the indicated key-value pairs as labels (it can have
                      additional labels as well). See https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector'
                    type: object
                  tolerations:
                    description: tolerations is a list of tolerations applied to the
                      relevant kind of pods See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
                      for more info. These are additional tolerations other than default
                      ones.
                    items:
                      description: The pod this Toleration is attached to tolerates
                        any taint that matches the triple <key,value,effect> using
                        the matching operator <operator>.
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match.
                            Empty means match all taint effects. When specified, allowed
                            values are NoSchedule, PreferNoSchedule and NoExecute.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies
                            to. Empty means match all taint keys. If the key is empty,
                            operator must be Exists; this combination means to match
                            all values and all keys.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to
                            the value. Valid operators are Exists and Equal. Defaults
                            to Equal. Exists is equivalent to wildcard for value,
                            so that a pod can tolerate all taints of a particular
                            category.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of
                            time the toleration (which must be of effect NoExecute,
                            otherwise this field is ignored) tolerates the taint.
                            By default, it is not set, which means tolerate the taint
                            forever (do not evict). Zero and negative values will
                            be treated as 0 (evict immediately) by the system.
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches
                            to. If the operator is Exists, the value should be empty,
                            otherwise just a regular string.
                          type: string
                      type: object
                    type: array
                type: object
              namespaceSelector:
                description: namespaces where pods should be gated before scheduling
                  Defaults to targeting namespaces with an "application-aware-quota/enable"
                  label key.
                properties:
                  matchExpressions:
                    description: matchExpressions is a list of label selector requirements.
                      The requirements are ANDed.
                    items:
                      description: A label selector requirement is a selector that
                        contains values, a key, and an operator that relates the key
                        and values.
                      properties:
                        key:
                          description: key is the label key that the selector applies
                            to.
                          type: string
                        operator:
                          description: operator represents a key's relationship to
                            a set of values. Valid operators are In, NotIn, Exists
                            and DoesNotExist.
                          type: string
                        values:
                          description: values is an array of string values. If the
                            operator is In or NotIn, the values array must be non-empty.
                            If the operator is Exists or DoesNotExist, the values
                            array must be empty. This array is replaced during a strategic
                            merge patch.
                          items:
                            type: string
                          type: array
                      required:
                      - key
                      - operator
                      type: object
                    type: array
                  matchLabels:
                    additionalProperties:
                      type: string
                    description: matchLabels is a map of {key,value} pairs. A single
                      {key,value} in the matchLabels map is equivalent to an element
                      of matchExpressions, whose key field is "key", the operator
                      is "In", and the values array contains only "value". The requirements
                      are ANDed.
                    type: object
                type: object
                x-kubernetes-map-type: atomic
              priorityClass:
                description: PriorityClass of the AAQ control plane
                type: string
              workload:
                description: Restrict on which nodes AAQ workload pods will be scheduled
                properties:
                  affinity:
                    description: affinity enables pod affinity/anti-affinity placement
                      expanding the types of constraints that can be expressed with
                      nodeSelector. affinity is going to be applied to the relevant
                      kind of pods in parallel with nodeSelector See https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for
                          the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node matches the corresponding matchExpressions;
                              the node(s) with the highest sum are the most preferred.
                            items:
                              description: An empty preferred scheduling term matches
                                all objects with implicit weight 0 (i.e. it's a no-op).
                                A null preferred scheduling term matches no objects
                                (i.e. is also a no-op).
                              properties:
                                preference:
                                  description: A node selector term, associated with
                                    the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                weight:
                                  description: Weight associated with matching the
                                    corresponding nodeSelectorTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to an update), the system
                              may or may not try to eventually evict the pod from
                              its node.
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms.
                                  The terms are ORed.
                                items:
                                  description: A null or empty node selector term
                                    matches no objects. The requirements of them are
                                    ANDed. The TopologySelectorTerm type implements
                                    a subset of the NodeSelectorTerm.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                            x-kubernetes-map-type: atomic
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                          co-locate this pod in the same node, zone, etc. as some
                          other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to a pod label update),
                              the system may or may not try to eventually evict the
                              pod from its node. When there are multiple elements,
                              the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules
                          (e.g. avoid putting this pod in the same node, zone, etc.
                          as some other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the anti-affinity expressions
                              specified by this field, but it may choose a node that
                              violates one or more of the expressions. The node that
                              is most preferred is the one with the greatest sum of
                              weights, i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              anti-affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified
                              by this field are not met at scheduling time, the pod
                              will not be scheduled onto the node. If the anti-affinity
                              requirements specified by this field cease to be met
                              at some point during pod execution (e.g. due to a pod
                              label update), the system may or may not try to eventually
                              evict the pod from its node. When there are multiple
                              elements, the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: 'nodeSelector is the node selector applied to the
                      relevant kind of pods It specifies a map of key-value pairs:
                      for the pod to be eligible to run on a node, the node must have
                      each of the indicated key-value pairs as labels (it can have
                      additional labels as well). See https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector'
                    type: object
                  tolerations:
                    description: tolerations is a list of tolerations applied to the
                      relevant kind of pods See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
                      for more info. These are additional tolerations other than default
                      ones.
                    items:
                      description: The pod this Toleration is attached to tolerates
                        any taint that matches the triple <key,value,effect> using
                        the matching operator <operator>.
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match.
                            Empty means match all taint effects. When specified, allowed
                            values are NoSchedule, PreferNoSchedule and NoExecute.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies
                            to. Empty means match all taint keys. If the key is empty,
                            operator must be Exists; this combination means to match
                            all values and all keys.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to
                            the value. Valid operators are Exists and Equal. Defaults
                            to Equal. Exists is equivalent to wildcard for value,
                            so that a pod can tolerate all taints of a particular
                            category.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of
                            time the toleration (which must be of effect NoExecute,
                            otherwise this field is ignored) tolerates the taint.
                            By default, it is not set, which means tolerate the taint
                            forever (do not evict). Zero and negative values will
                            be treated as 0 (evict immediately) by the system.
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches
                            to. If the operator is Exists, the value should be empty,
                            otherwise just a regular string.
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: AAQStatus defines the status of the installation
            properties:
              conditions:
                description: A list of current conditions of the resource
                items:
                  description: Condition represents the state of the operator's reconciliation
                    functionality.
                  properties:
                    lastHeartbeatTime:
                      format: date-time
                      type: string
                    lastTransitionTime:
                      format: date-time
                      type: string
                    message:
                      type: string
                    reason:
                      type: string
                    status:
                      type: string
                    type:
                      description: ConditionType is the state of the operator's reconciliation
                        functionality.
                      type: string
                  required:
                  - status
                  - type
                  type: object
                type: array
              observedVersion:
                description: The observed version of the resource
                type: string
              operatorVersion:
                description: The version of the resource as defined by the operator
                type: string
              phase:
                description: Phase is the current phase of the deployment
                type: string
              targetVersion:
                description: The desired version of the resource
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"aaqjobqueueconfig": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: aaqjobqueueconfigs.aaq.kubevirt.io
spec:
  group: aaq.kubevirt.io
  names:
    categories:
    - all
    kind: AAQJobQueueConfig
    listKind: AAQJobQueueConfigList
    plural: aaqjobqueueconfigs
    shortNames:
    - aaqjqc
    - aaqjqcs
    singular: aaqjobqueueconfig
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: AAQJobQueueConfig
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: CA configuration CA certs are kept in the CA bundle as long
              as they are valid
            type: object
          status:
            description: AAQJobQueueConfigStatus defines the status with metadata
              for current jobs
            properties:
              controllerLock:
                additionalProperties:
                  type: boolean
                type: object
              podsInJobQueue:
                description: BuiltInCalculationConfigToApply
                items:
                  type: string
                type: array
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"applicationawareappliedclusterresourcequota": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: applicationawareappliedclusterresourcequotas.aaq.kubevirt.io
spec:
  group: aaq.kubevirt.io
  names:
    categories:
    - all
    kind: ApplicationAwareAppliedClusterResourceQuota
    listKind: ApplicationAwareAppliedClusterResourceQuotaList
    plural: applicationawareappliedclusterresourcequotas
    shortNames:
    - aacrq
    - aacrqs
    singular: applicationawareappliedclusterresourcequota
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ApplicationAwareAppliedClusterResourceQuota mirrors ApplicationAwareAppliedClusterResourceQuota
          at a project scope, for projection into a project.  It allows a project-admin
          to know which ClusterResourceQuotas are applied to his project and their
          associated usage.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: Spec defines the desired quota
            properties:
              quota:
                description: Quota defines the desired quota
                properties:
                  hard:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: 'hard is the set of desired hard limits for each
                      named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/'
                    type: object
                  scopeSelector:
                    description: scopeSelector is also a collection of filters like
                      scopes that must match each object tracked by a quota but expressed
                      using ScopeSelectorOperator in combination with possible values.
                      For a resource to match, both scopes AND scopeSelector (if specified
                      in spec), must be matched.
                    properties:
                      matchExpressions:
                        description: A list of scope selector requirements by scope
                          of the resources.
                        items:
                          description: A scoped-resource selector requirement is a
                            selector that contains values, a scope name, and an operator
                            that relates the scope name and values.
                          properties:
                            operator:
                              description: Represents a scope's relationship to a
                                set of values. Valid operators are In, NotIn, Exists,
                                DoesNotExist.
                              type: string
                            scopeName:
                              description: The name of the scope that the selector
                                applies to.
                              type: string
                            values:
                              description: An array of string values. If the operator
                                is In or NotIn, the values array must be non-empty.
                                If the operator is Exists or DoesNotExist, the values
                                array must be empty. This array is replaced during
                                a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - operator
                          - scopeName
                          type: object
                        type: array
                    type: object
                    x-kubernetes-map-type: atomic
                  scopes:
                    description: A collection of filters that must match each object
                      tracked by a quota. If not specified, the quota matches all
                      objects.
                    items:
                      description: A ResourceQuotaScope defines a filter that must
                        match each object tracked by a quota
                      type: string
                    type: array
                type: object
              selector:
                description: Selector is the selector used to match projects. It should
                  only select active projects on the scale of dozens (though it can
                  select many more less active projects).  These projects will contend
                  on object creation through this resource.
                properties:
                  annotations:
                    additionalProperties:
                      type: string
                    description: AnnotationSelector is used to select projects by
                      annotation.
                    nullable: true
                    type: object
                  labels:
                    description: LabelSelector is used to select projects by label.
                    nullable: true
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                    x-kubernetes-map-type: atomic
                type: object
            required:
            - quota
            - selector
            type: object
          status:
            description: Status defines the actual enforced quota and its current
              usage
            properties:
              namespaces:
                description: Namespaces slices the usage by project.  This division
                  allows for quick resolution of deletion reconciliation inside of
                  a single project without requiring a recalculation across all projects.  This
                  can be used to pull the deltas for a given project.
                items:
                  description: ResourceQuotaStatusByNamespace gives status for a particular
                    project
                  properties:
                    namespace:
                      description: Namespace the project this status applies to
                      type: string
                    status:
                      description: Status indicates how many resources have been consumed
                        by this project
                      properties:
                        hard:
                          additionalProperties:
                            anyOf:
                            - type: integer
                            - type: string
                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                            x-kubernetes-int-or-string: true
                          description: 'Hard is the set of enforced hard limits for
                            each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/'
                          type: object
                        used:
                          additionalProperties:
                            anyOf:
                            - type: integer
                            - type: string
                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                            x-kubernetes-int-or-string: true
                          description: Used is the current observed total usage of
                            the resource in the namespace.
                          type: object
                      type: object
                  required:
                  - namespace
                  - status
                  type: object
                nullable: true
                type: array
              total:
                description: Total defines the actual enforced quota and its current
                  usage across all projects
                properties:
                  hard:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: 'Hard is the set of enforced hard limits for each
                      named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/'
                    type: object
                  used:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: Used is the current observed total usage of the resource
                      in the namespace.
                    type: object
                type: object
            required:
            - total
            type: object
        required:
        - metadata
        - spec
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"applicationawareclusterresourcequota": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: applicationawareclusterresourcequotas.aaq.kubevirt.io
spec:
  group: aaq.kubevirt.io
  names:
    kind: ApplicationAwareClusterResourceQuota
    listKind: ApplicationAwareClusterResourceQuotaList
    plural: applicationawareclusterresourcequotas
    shortNames:
    - acrq
    - acrqs
    singular: applicationawareclusterresourcequota
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ApplicationAwareClusterResourceQuota mirrors ResourceQuota at
          a cluster scope.  This object is easily convertible to synthetic ResourceQuota
          object to allow quota evaluation re-use.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: Spec defines the desired quota
            properties:
              quota:
                description: Quota defines the desired quota
                properties:
                  hard:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: 'hard is the set of desired hard limits for each
                      named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/'
                    type: object
                  scopeSelector:
                    description: scopeSelector is also a collection of filters like
                      scopes that must match each object tracked by a quota but expressed
                      using ScopeSelectorOperator in combination with possible values.
                      For a resource to match, both scopes AND scopeSelector (if specified
                      in spec), must be matched.
                    properties:
                      matchExpressions:
                        description: A list of scope selector requirements by scope
                          of the resources.
                        items:
                          description: A scoped-resource selector requirement is a
                            selector that contains values, a scope name, and an operator
                            that relates the scope name and values.
                          properties:
                            operator:
                              description: Represents a scope's relationship to a
                                set of values. Valid operators are In, NotIn, Exists,
                                DoesNotExist.
                              type: string
                            scopeName:
                              description: The name of the scope that the selector
                                applies to.
                              type: string
                            values:
                              description: An array of string values. If the operator
                                is In or NotIn, the values array must be non-empty.
                                If the operator is Exists or DoesNotExist, the values
                                array must be empty. This array is replaced during
                                a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - operator
                          - scopeName
                          type: object
                        type: array
                    type: object
                    x-kubernetes-map-type: atomic
                  scopes:
                    description: A collection of filters that must match each object
                      tracked by a quota. If not specified, the quota matches all
                      objects.
                    items:
                      description: A ResourceQuotaScope defines a filter that must
                        match each object tracked by a quota
                      type: string
                    type: array
                type: object
              selector:
                description: Selector is the selector used to match projects. It should
                  only select active projects on the scale of dozens (though it can
                  select many more less active projects).  These projects will contend
                  on object creation through this resource.
                properties:
                  annotations:
                    additionalProperties:
                      type: string
                    description: AnnotationSelector is used to select projects by
                      annotation.
                    nullable: true
                    type: object
                  labels:
                    description: LabelSelector is used to select projects by label.
                    nullable: true
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                    x-kubernetes-map-type: atomic
                type: object
            required:
            - quota
            - selector
            type: object
          status:
            description: Status defines the actual enforced quota and its current
              usage
            properties:
              namespaces:
                description: Namespaces slices the usage by project.  This division
                  allows for quick resolution of deletion reconciliation inside of
                  a single project without requiring a recalculation across all projects.  This
                  can be used to pull the deltas for a given project.
                items:
                  description: ResourceQuotaStatusByNamespace gives status for a particular
                    project
                  properties:
                    namespace:
                      description: Namespace the project this status applies to
                      type: string
                    status:
                      description: Status indicates how many resources have been consumed
                        by this project
                      properties:
                        hard:
                          additionalProperties:
                            anyOf:
                            - type: integer
                            - type: string
                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                            x-kubernetes-int-or-string: true
                          description: 'Hard is the set of enforced hard limits for
                            each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/'
                          type: object
                        used:
                          additionalProperties:
                            anyOf:
                            - type: integer
                            - type: string
                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                            x-kubernetes-int-or-string: true
                          description: Used is the current observed total usage of
                            the resource in the namespace.
                          type: object
                      type: object
                  required:
                  - namespace
                  - status
                  type: object
                nullable: true
                type: array
              total:
                description: Total defines the actual enforced quota and its current
                  usage across all projects
                properties:
                  hard:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: 'Hard is the set of enforced hard limits for each
                      named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/'
                    type: object
                  used:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: Used is the current observed total usage of the resource
                      in the namespace.
                    type: object
                type: object
            required:
            - total
            type: object
        required:
        - metadata
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"applicationawareresourcequota": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: applicationawareresourcequotas.aaq.kubevirt.io
spec:
  group: aaq.kubevirt.io
  names:
    categories:
    - all
    kind: ApplicationAwareResourceQuota
    listKind: ApplicationAwareResourceQuotaList
    plural: applicationawareresourcequotas
    shortNames:
    - arq
    - arqs
    singular: applicationawareresourcequota
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ApplicationAwareResourceQuota defines resources that should be
          reserved for a VMI migration
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ApplicationAwareResourceQuotaSpec is an extension of corev1.ResourceQuotaSpec
            properties:
              hard:
                additionalProperties:
                  anyOf:
                  - type: integer
                  - type: string
                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                  x-kubernetes-int-or-string: true
                description: 'hard is the set of desired hard limits for each named
                  resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/'
                type: object
              scopeSelector:
                description: scopeSelector is also a collection of filters like scopes
                  that must match each object tracked by a quota but expressed using
                  ScopeSelectorOperator in combination with possible values. For a
                  resource to match, both scopes AND scopeSelector (if specified in
                  spec), must be matched.
                properties:
                  matchExpressions:
                    description: A list of scope selector requirements by scope of
                      the resources.
                    items:
                      description: A scoped-resource selector requirement is a selector
                        that contains values, a scope name, and an operator that relates
                        the scope name and values.
                      properties:
                        operator:
                          description: Represents a scope's relationship to a set
                            of values. Valid operators are In, NotIn, Exists, DoesNotExist.
                          type: string
                        scopeName:
                          description: The name of the scope that the selector applies
                            to.
                          type: string
                        values:
                          description: An array of string values. If the operator
                            is In or NotIn, the values array must be non-empty. If
                            the operator is Exists or DoesNotExist, the values array
                            must be empty. This array is replaced during a strategic
                            merge patch.
                          items:
                            type: string
                          type: array
                      required:
                      - operator
                      - scopeName
                      type: object
                    type: array
                type: object
                x-kubernetes-map-type: atomic
              scopes:
                description: A collection of filters that must match each object tracked
                  by a quota. If not specified, the quota matches all objects.
                items:
                  description: A ResourceQuotaScope defines a filter that must match
                    each object tracked by a quota
                  type: string
                type: array
            type: object
          status:
            description: ApplicationAwareResourceQuotaStatus is an extension of corev1.ResourceQuotaStatus
            properties:
              hard:
                additionalProperties:
                  anyOf:
                  - type: integer
                  - type: string
                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                  x-kubernetes-int-or-string: true
                description: 'Hard is the set of enforced hard limits for each named
                  resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/'
                type: object
              used:
                additionalProperties:
                  anyOf:
                  - type: integer
                  - type: string
                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                  x-kubernetes-int-or-string: true
                description: Used is the current observed total usage of the resource
                  in the namespace.
                type: object
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
}
