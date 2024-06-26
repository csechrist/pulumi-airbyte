// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * DestinationDuckdb DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const myDestinationDuckdb = airbyte.getDestinationDuckdb({
 *     destinationId: "...my_destination_id...",
 * });
 * ```
 */
export function getDestinationDuckdb(args: GetDestinationDuckdbArgs, opts?: pulumi.InvokeOptions): Promise<GetDestinationDuckdbResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getDestinationDuckdb:getDestinationDuckdb", {
        "destinationId": args.destinationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDestinationDuckdb.
 */
export interface GetDestinationDuckdbArgs {
    destinationId: string;
}

/**
 * A collection of values returned by getDestinationDuckdb.
 */
export interface GetDestinationDuckdbResult {
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
 * DestinationDuckdb DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const myDestinationDuckdb = airbyte.getDestinationDuckdb({
 *     destinationId: "...my_destination_id...",
 * });
 * ```
 */
export function getDestinationDuckdbOutput(args: GetDestinationDuckdbOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDestinationDuckdbResult> {
    return pulumi.output(args).apply((a: any) => getDestinationDuckdb(a, opts))
}

/**
 * A collection of arguments for invoking getDestinationDuckdb.
 */
export interface GetDestinationDuckdbOutputArgs {
    destinationId: pulumi.Input<string>;
}
