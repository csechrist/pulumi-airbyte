// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * DestinationMongodb DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const myDestinationMongodb = airbyte.getDestinationMongodb({
 *     destinationId: "...my_destination_id...",
 * });
 * ```
 */
export function getDestinationMongodb(args: GetDestinationMongodbArgs, opts?: pulumi.InvokeOptions): Promise<GetDestinationMongodbResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getDestinationMongodb:getDestinationMongodb", {
        "destinationId": args.destinationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDestinationMongodb.
 */
export interface GetDestinationMongodbArgs {
    destinationId: string;
}

/**
 * A collection of values returned by getDestinationMongodb.
 */
export interface GetDestinationMongodbResult {
    /**
     * The values required to configure the destination. Parsed as JSON.
     */
    readonly configuration: string;
    readonly destinationId: string;
    readonly destinationType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly workspaceId: string;
}
/**
 * DestinationMongodb DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const myDestinationMongodb = airbyte.getDestinationMongodb({
 *     destinationId: "...my_destination_id...",
 * });
 * ```
 */
export function getDestinationMongodbOutput(args: GetDestinationMongodbOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDestinationMongodbResult> {
    return pulumi.output(args).apply((a: any) => getDestinationMongodb(a, opts))
}

/**
 * A collection of arguments for invoking getDestinationMongodb.
 */
export interface GetDestinationMongodbOutputArgs {
    destinationId: pulumi.Input<string>;
}
